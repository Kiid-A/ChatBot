package main

import ()"fmt"

func main() {
	err := chatBot.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
