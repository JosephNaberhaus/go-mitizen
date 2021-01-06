package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetExecPath() (path string, err error) {
	cmd := exec.Command("git", "--exec-path")

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting git exec path: %w", err)
	}

	return strings.ReplaceAll(string(output), "\n", ""), nil
}

func GetRepositoryRoot() (path string, err error) {
	cmd := exec.Command("git",  "rev-parse", "--show-toplevel")

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting repository root: %w", err)
	}

	return strings.ReplaceAll(string(output), "\n", ""), nil
}

func IsInGitRepository() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")

	_, err := cmd.Output()
	return err == nil
}