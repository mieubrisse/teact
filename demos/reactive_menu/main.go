package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/bubblebath"
	"github.com/mieubrisse/teact/demos/reactive_menu/app"
	"os"
)

func main() {
	myApp := app.New()

	if _, err := bubblebath.RunBubbleBathProgram(
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
