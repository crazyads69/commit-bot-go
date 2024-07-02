package utils

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"commit-bot/schemas" // Assuming your schemas package is in the correct path

	"github.com/google/generative-ai-go/genai"
)

// GenerateCommitMessage generates a commit message using the Gemini model.
func GenerateCommitMessage(diff string, model *genai.GenerativeModel, ctx context.Context) string {
	prompt := SYSTEM_PROMPT + diff + COMMIT_MSG_STRUCTURE

	// Generate commit message
	response, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal("Error generating commit message:", err)
	}
	// Get the generated commit message
	responseJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatal("Error marshalling response to JSON:", err)
	}
	// Parse the JSON response
	var responseData schemas.Response
	if err := json.Unmarshal([]byte(responseJSON), &responseData); err != nil {
		log.Fatal("Error parsing JSON response:", err)
	}

	// Extract and join the commit message parts
	if len(responseData.Candidates) > 0 {
		return strings.Join(responseData.Candidates[0].Content.Parts, "")
	}

	// Handle the case where no candidates are returned
	log.Println("Warning: No commit message candidates generated. Returning an empty message.")
	return ""
}
