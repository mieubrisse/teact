package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	app "github.com/mieubrisse/teact/demos/reactive_form/app"
	"github.com/mieubrisse/teact/teact"
	"os"
)

func main() {
	if _, err := teact.RunTeact(app.New(), tea.WithAltScreen()); err != nil {
		fmt.Printf("An error occurred running the program:\n%v", err)
		os.Exit(1)
	}
}
