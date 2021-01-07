package git

import "os/exec"

func AreStagedFiles() bool {
	cmd := exec.Command("git", "diff", "--quiet", "--cached", "--exit-code")

	_, err := cmd.Output()
	return err != nil
}
