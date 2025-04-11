package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/LeonibelDev/gemini-api-console/internal/service"
	"github.com/joho/godotenv"
)

func chat(reader *bufio.Reader) {
	var message string

	// Get user input
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	message = strings.TrimSpace(message)

	// Check if the user wants to exit
	if message == "exit" {
		os.Exit(0)
	}

	// Manage response
	response, err := service.Gemini(message)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", response)
	}

}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load("config/.env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		chat(reader)
	}

}
