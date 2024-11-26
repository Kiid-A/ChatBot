package cmd

import (
	"ChatBot/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

func AskBot(req model.ChatRequest) (model.ChatReply, error) {
	jsonValue, err := json.Marshal(req)
	if err != nil {
		return , err
	}

	resp, err := http.Post("http://localhost:5000/chat", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var chatResp model.ChatReply
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return "", err
	}

}
