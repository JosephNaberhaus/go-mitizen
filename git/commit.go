package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Commit(message string) error {
	escaped := strings.ReplaceAll(message, "\"", "\\\"")

	cmd := exec.Command("git", "commit", "-m", escaped)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error commiting: %w", err)
	}

	return nil
}
