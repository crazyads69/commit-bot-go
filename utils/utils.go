package utils

import "strings"

// Define the clean function that removes special characters from the commit message
func CleanSpecialCharacter(commitMessage string) string {
	// Define the special characters to remove (Markdown syntax)
	specialCharacters := []string{"```", "**", "`", "#", "*"}
	// Remove the special characters
	for _, character := range specialCharacters {
		commitMessage = strings.ReplaceAll(commitMessage, character, "")
	}
	// Strip the commit message
	commitMessage = strings.TrimSpace(commitMessage)
	return commitMessage
}

// Define the function that validates the commit message structure
func ValidateCommitMessage(commitMessage string) bool {
	// Define the commit message structure
	// <type>[optional scope]: <short summary>
	// Split the commit message into parts
	parts := strings.Split(commitMessage, ":")
	// Get the type of the commit message
	commitType := strings.Split(parts[0], "(")[0]
	// Check if the commit type is valid
	if commitType != "feat" && commitType != "fix" && commitType != "docs" && commitType != "style" && commitType != "refactor" && commitType != "perf" && commitType != "test" && commitType != "chore" {
		return false
	}
	return true
}
