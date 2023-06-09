package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/demos/hello_world/greeter"
	"github.com/mieubrisse/teact/teact"
	"os"
)

func main() {
	myApp := greeter.New()
	if _, err := teact.Run(myApp, tea.WithAltScreen()); err != nil {
		fmt.Printf("An error occurred running the program:\n%v", err)
		os.Exit(1)
	}
}
