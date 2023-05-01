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

// Utility for dealing with tea.Msg events
// Returns an emptystring if the object isn't a tea.KeyMsg
func GetMaybeKeyMsgStr(msg tea.Msg) string {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return msg.String()
	}
	return ""
}
