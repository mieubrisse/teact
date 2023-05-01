package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/bubblebath"
	"github.com/mieubrisse/teact/demos/journal/app"
	"os"
)

func main() {
	myApp := app.New()
	myApp.SetFocus(true)

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
