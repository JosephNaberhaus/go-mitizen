package git

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func InstallAsSubcommand(subcommandName string) error {
	execPath := os.Args[0]

	gitExecPath, err := GetExecPath()
	if err != nil {
		return err
	}

	execContent, err := ioutil.ReadFile(execPath)
	if err != nil {
		return fmt.Errorf("error reading executable to install: %w", err)
	}

	installPath := filepath.Join(gitExecPath, "git-" +subcommandName)
	err = ioutil.WriteFile(installPath, execContent, 0x755)
	if err != nil {
		return fmt.Errorf("error writing executable subcommand: %w", err)
	}

	return nil
}
