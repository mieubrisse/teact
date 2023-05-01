package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/demos/reactive_form/secret_agent_terminal"
	"github.com/mieubrisse/teact/teact"
	"os"
)

func main() {
	if _, err := teact.RunTeact(secret_agent_terminal.New(), tea.WithAltScreen()); err != nil {
		fmt.Printf("An error occurred running the program:\n%v", err)
		os.Exit(1)
	}
}
