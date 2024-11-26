package cmd

import "os/exec"

type SysCmd interface {
	RunBot() error
	StopBot() error
	AskBot() (string, error)
}

func RunBot() error {
	if err := exec.Command("python", "cmd/test"); err != nil {
		return fmt.
	}

	return nil
}
