package cmd

import "os/exec"

type SysCmd interface {
	RunBot() error
	StopBot() error
	AskBot() (string, error
}

func RunBot() error {
	exec.Command("python", "cmd/test")
}
