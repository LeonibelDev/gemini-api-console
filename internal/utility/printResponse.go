package utility

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
)

func PrintResponse(markdown string) {
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
