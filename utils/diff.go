package utils

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// GetDiffChanges retrieves and formats git diff changes with committer info.
func GetDiffChanges() string {
	// Stage all changes (using combined output for brevity)
	if err := exec.Command("git", "add", ".").Run(); err != nil {
		log.Fatal("Error staging changes:", err)
	}

	// Get staged diff
	diffOutput, err := exec.Command("git", "diff", "--staged").Output()
	if err != nil {
		log.Fatal("Error getting diff:", err)
	}

	// Get committer info
	committerOutput, err := exec.Command("git", "var", "GIT_COMMITTER_IDENT").Output()
	if err != nil {
		log.Fatal("Error getting committer info:", err)
	}
	committer := string(bytes.Split(committerOutput, []byte(">"))[0]) // Extract before '>'

	log.Println("Found diff changes and committer info.")
	log.Printf("Committer: %s>\n", committer)

	// Combine diff and committer info efficiently
	return fmt.Sprintf("%s\n\nCommitter: %s>", string(diffOutput), committer)
}
