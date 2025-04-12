package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"

	"github.com/LeonibelDev/gemini-api-console/internal/service"
	"github.com/joho/godotenv"
)

var chatHistory []string

func chat(reader *bufio.Reader) {
	chatHistory = append(chatHistory, "System: You are a helpful AI assistant. Keep track of the conversation and respond accordingly.")

	// Get user input
	fmt.Print("> ")
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
		os.Exit(0)
	}

	// Create a promp
	fullPromp := strings.Join(chatHistory, "\n")

	// Manage response
	response, err := service.Gemini(fullPromp)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		chatHistory = append(chatHistory, "Gemini: "+response)
		printResponse(response)
	}

}

func printResponse(markdown string) {
	// Create a new renderer with appropriate options
	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(100),
	)
	if err != nil {
		fmt.Println("Error creating renderer:", err)
		return
	}

	// Render the markdown
	rendered, err := renderer.Render(markdown)
	if err != nil {
		fmt.Println("Error rendering markdown:", err)
		return
	}

	// Split the rendered text into lines and print them one by one
	for _, line := range strings.Split(rendered, "\n") {
		fmt.Println(line)
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println()
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load("config/.env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Gemini API Console! Type 'exit' to quit.")

	for {
		chat(reader)
	}

}
