package utils

import (
	"commit-bot/schemas"
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

func GenerateCommitMessage(diff string, model *genai.GenerativeModel, ctx context.Context) string {
	prompt := SYSTEM_PROMPT + diff + COMMIT_MSG_STRUCTURE
	// Generate a commit message using the model
	commitMessage, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}
	// Get the generated commit message
	commitMessageJSON, err := json.MarshalIndent(commitMessage, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	// Unmarshal the JSON string into the struct
	var response schemas.Response
	err = json.Unmarshal([]byte(commitMessageJSON), &response)
	if err != nil {
		panic(err)
	}

	// Access the "Parts" array
	parts := response.Candidates[0].Content.Parts
	// Join the parts to form the commit message
	commitMsg := []byte(strings.Join(parts, ""))
	return string(commitMsg)
}
