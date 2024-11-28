package cmd

import (
	"ChatBot/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

var (
	pid int
)

type SysCmd interface {
	RunBot() error
	EndBot() error
	AskBot(string) (string, error)
}

func RunBot() error {
	cmd := exec.Command("python3", "cmd/chatbot.py")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start ChatBot: %w", err)
	}
	pid = cmd.Process.Pid
	return nil
}

func EndBot() error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("failed to find pid %d %w", pid, err)
	}

	if err := process.Kill(); err != nil {
		return fmt.Errorf("failed to kill pid %d %w", pid, err)
	}

	return nil
}

func AskBot(req model.ChatRequest) (model.ChatReply, error) {
	jsonValue, err := json.Marshal(req)
	if err != nil {
		return model.ChatReply{}, err
	}

	fmt.Println(bytes.NewBuffer(jsonValue))

	resp, err := http.Post("http://localhost:5000/chat/"+req.UserId,
		"application/json", bytes.NewBuffer(jsonValue))
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
