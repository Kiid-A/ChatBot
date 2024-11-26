package cmd

import (
	"fmt"
	"os/exec"
)

type SysCmd interface {
	RunBot() error
	StopBot() error
	AskBot(string) (string, error)
}

func RunBot() error {
	cmd := exec.Command("python", "cmd/test")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start ChatBot: %w", err)
	}
	return nil
}

func AskBot()
