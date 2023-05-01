package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/demos/journal/app"
	"github.com/mieubrisse/teact/teact"
	"os"
)

func main() {
	myApp := app.New()
	myApp.SetFocus(true)

	if _, err := teact.Run(myApp, tea.WithAltScreen()); err != nil {
		fmt.Printf("An error occurred running the program:\n%v", err)
		os.Exit(1)
	}
}
