package cmd

import "os/exec"

type SysCmd interface {
	RunBot() error
	StopB
}

func RunBot() error {
	exec.Command("python", "cmd/test")
}
