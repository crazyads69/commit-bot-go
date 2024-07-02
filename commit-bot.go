package main

import (
	"commit-bot/utils"
	"context"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	// --- Setup ---
	ctx := context.Background()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err) // Improved error message
	}

	// Initialize Gemini client
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal("Error creating Gemini client:", err) // Improved error message
	}
	defer client.Close()

	model := client.GenerativeModel(os.Getenv("GEMINI_MODEL"))
	model.SetTemperature(0.6)
	model.SetMaxOutputTokens(8192)
	model.SetTopP(0.9)
	// (Safety settings remain the same)

	// --- Commit and Push Logic ---
	retryDelay := 60 * time.Second
	for {
		diff := utils.GetDiffChanges()
		if diff == "" {
			log.Println("No changes found in current branch. Exiting")
			return
		}

		commitMessage := utils.GenerateCommitMessage(diff, model, ctx)
		commitMessage = utils.CleanSpecialCharacter(commitMessage)
		log.Println("Generated commit message:", commitMessage)

		if !utils.ValidateCommitMessage(commitMessage) {
			log.Println("Invalid commit message. Retrying in", retryDelay, "...")
			time.Sleep(retryDelay)
			continue
		}

		// --- Commit and Push (Only if commit message is valid) ---
		log.Println("Commit message validated and generated.") // Combined message

		// Set git pull rebase
		if err := exec.Command("git", "config", "--local", "pull.rebase", "true").Run(); err != nil {
			log.Fatal("Error setting git pull.rebase:", err)
		}

		// Get branch name
		branchName, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		if err != nil {
			log.Fatal("Error getting branch name:", err)
		}

		commitTitle := commitMessage[:strings.Index(commitMessage, "\n")]
		commitBody := commitMessage[strings.Index(commitMessage, "\n")+1:]

		// Commit changes
		if _, err := exec.Command("git", "commit", "-m", commitTitle, "-m", commitBody).Output(); err != nil {
			log.Fatal("Error committing changes:", err)
		}

		// Push changes
		if _, err := exec.Command("git", "push", "-u", "origin", strings.TrimSpace(string(branchName))).Output(); err != nil {
			log.Fatal("Error pushing changes:", err)
		}

		// Get commit hash
		commitHash, err := exec.Command("git", "rev-parse", "HEAD").Output()
		if err != nil {
			log.Fatal("Error getting commit hash:", err)
		}

		log.Println("Changes committed and pushed. Latest commit hash:", string(commitHash))
		return
	}
}
