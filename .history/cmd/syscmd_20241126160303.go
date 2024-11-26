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
	EndBot() error
	AskBot(string) (string, error)
}

func RunBot() error {
	cmd := exec.Command("python", "./test.py")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start ChatBot: %w", err)
	}
	return nil
}

func EndBot() error {
	return nil
}

func AskBot(req model.ChatRequest) (model.ChatReply, error) {
	jsonValue, err := json.Marshal(req)
	if err != nil {
		return model.ChatReply{}, err
	}

	resp, err := http.Post("http://localhost:5000/chat", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return model.ChatReply{}, err
	}
	defer resp.Body.Close()

	var chatResp model.ChatReply
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return model.ChatReply{}, err
	}

	return chatResp, nil
}
