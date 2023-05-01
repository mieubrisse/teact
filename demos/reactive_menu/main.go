package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/demos/reactive_menu/app"
	"github.com/mieubrisse/teact/teact"
	"os"
)

func main() {
	myApp := app.New()

	if _, err := teact.RunTeactApp(
		myApp,
		nil,
		[]tea.ProgramOption{
			tea.WithAltScreen(),
		},
	); err != nil {
		fmt.Printf("An error occurred running the program:\n%v", err)
		os.Exit(1)
	}
}
