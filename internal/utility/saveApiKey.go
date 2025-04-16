package utility

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SaveAPIKEY() {
	if key := os.Getenv("GEMINI_API_KEY"); key != "" {
		return
	}

	fmt.Print("Type your GEMINI API KEY: ")

	reader := bufio.NewReader(os.Stdin)
	apiKey, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(0)
	}

	apiKey = strings.TrimSpace(apiKey)

	if apiKey == "" {
		fmt.Println("API KEY cannot be empty")
		os.Exit(0)
	}

	if err := os.Setenv("GEMINI_API_KEY", apiKey); err != nil {
		fmt.Println("Error to set env:", err)
	}

}
