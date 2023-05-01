package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/demos/scoreboard/scoreboard"
	"github.com/mieubrisse/teact/teact"
	"os"
)

func main() {
	// Can configure the scoreboard
	myApp := scoreboard.New(
		scoreboard.WithHomeScore(3),
		scoreboard.WithAwayScore(2),
	)
	if _, err := teact.Run(myApp, tea.WithAltScreen()); err != nil {
		fmt.Printf("An error occurred running the program:\n%v", err)
		os.Exit(1)
	}
}
