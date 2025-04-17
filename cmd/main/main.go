package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/LeonibelDev/gemini-api-console/internal/utility"
	"github.com/joho/godotenv"
)

func main() {
	// Flags config
	utility.Flags()

	// Clear console
	utility.ClearConsole()

	// Load environment variables from .env file
	if os.Getenv("GEMINI_API_KEY") == "" {
		if err := godotenv.Load("../../config/.env"); err != nil {
			fmt.Println("Error loading .env file")
			os.Exit(0)
		}
	}

	utility.SaveAPIKEY()

	// User funtionality
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Gemini API Console! Type 'exit' to quit.")

	for {
		utility.Chat(reader)
	}

}
