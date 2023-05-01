package utilities

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/teact/components"
)

// Tries to pass a tea.Msg to a component, doing nothing if it's not interactive
// This is useful when we're not sure if we'll get an InteractiveComponent or not
func TryUpdate(component components.Component, msg tea.Msg) tea.Cmd {
	switch castedComponent := component.(type) {
	case components.InteractiveComponent:
		return castedComponent.Update(msg)
	}
	return nil
}
