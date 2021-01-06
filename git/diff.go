package git

import "os/exec"

func AreStagedFiles() bool {
	cmd := exec.Command("git", "diff", "--cached", "--exit-code")

	_, err := cmd.Output()
	return err != nil
}
