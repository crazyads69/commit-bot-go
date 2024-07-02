package utils

import (
	"regexp"
	"strings"
)

// CleanSpecialCharacter removes unwanted characters from a commit message.
func CleanSpecialCharacter(commitMessage string) string {
	// Use a regular expression for more flexible pattern matching
	re := regexp.MustCompile("```|\\*\\*|`|#|\\*")
	commitMessage = re.ReplaceAllString(commitMessage, "")

	return strings.TrimSpace(commitMessage)
}

// ValidateCommitMessage checks if a commit message follows the conventional format.
func ValidateCommitMessage(commitMessage string) bool {
	// More efficient and readable regular expression for validation
	validCommit := regexp.MustCompile(`^(feat|fix|docs|style|refactor|perf|test|chore|ci|build|revert)(?:\(.+?\))?\:?\s.+`)
	return validCommit.MatchString(commitMessage)
}
