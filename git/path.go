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
