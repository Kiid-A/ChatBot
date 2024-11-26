package cmd

import (
	"ChatBot/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
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

func AskBot(req model.ChatRequest) (string, error) {
	jsonValue, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://localhost:5000/chat", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()

	// 读取Flask应用的响应
	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response from Flask app"})
		return
	}

}
