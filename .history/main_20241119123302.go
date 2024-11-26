package main

import (
	"fmt"
	"ChatBot/cmd/chatbot"
)

func main() {
	err := chatBot.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
