package utility

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/LeonibelDev/gemini-api-console/internal/service"
)

var chatHistory []string

func Chat(reader *bufio.Reader) {
	chatHistory = append(chatHistory, "System: You are a helpful AI assistant. Keep track of the conversation and respond accordingly. Verify the user language and response with the same language, dont say your role like gemini: ...")

	// Get user input
	fmt.Print("➜ ")
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	message = strings.TrimSpace(message)

	// Add user input to history
	chatHistory = append(chatHistory, "User: "+message)

	// Check if the user wants to exit
	if strings.ToLower(message) == "exit" {
		ClearConsole()
		os.Exit(0)
	}

	// Create a promp
	fullPromp := strings.Join(chatHistory, "\n")

	// Manage response
	response, err := service.Gemini(fullPromp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	chatHistory = append(chatHistory, "Gemini: "+response)
	PrintResponse(response)

}
