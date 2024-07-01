package main

import (
	"commit-bot/utils"
	"context"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	// Load the API key from the environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// The Gemini 1.5 models are versatile and work with most use cases
	model := client.GenerativeModel(os.Getenv("GEMINI_MODEL"))
	// Configure the model
	model.SetTemperature(0.6)
	model.SetMaxOutputTokens(8192)
	model.SetTopP(0.9)
	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockNone,
		},
	}

	// Get the diff changes
	diff := utils.GetDiffChanges()
	if diff == "" {
		// If there are no changes, exit
		log.Println("No changes found")
		log.Println("Exiting")
		return
	}
	// Commit the changes
	// Generate a commit message
	commitMessage := utils.GenerateCommitMessage(diff, model, ctx)
	// Remove the special characters from the commit message
	commitMessage = utils.CleanSpecialCharacter(commitMessage)
	// Validate the commit message that ensure it is following the commit message structure guidelines <type>[optional scope]: <short summary>
	if !utils.ValidateCommitMessage(commitMessage) {
		log.Println("Invalid commit message")
		log.Println("Exiting")
		return
	}
	log.Println("Commit message validated")
	log.Println("Commit message generated")
	// Get current branch name to commit the changes
	branchName, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Branch name: " + string(branchName))
	// Get the first line of the commit message as the commit title
	commitTitle := commitMessage[:strings.Index(commitMessage, "\n")]
	// Get the leftover part of the commit message as the commit body
	commitBody := commitMessage[strings.Index(commitMessage, "\n")+1:]
	// Commit the changes with the generated commit message
	_, err = exec.Command("git", "commit", "-m", commitTitle, "-m", commitBody).Output()
	if err != nil {
		log.Fatal(err)
	}
	// Push the changes to the current branch
	_, err = exec.Command("git", "push", "-u", "origin", string(branchName)).Output()
	if err != nil {
		log.Fatal(err)
	}
}
