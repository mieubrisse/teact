package components

import tea "github.com/charmbracelet/bubbletea"

// Tries to pass a tea.Msg to a component, doing nothing if it's not interactive
// This is useful when we're not sure if we'll get an InteractiveComponent or not
func TryUpdate(component Component, msg tea.Msg) tea.Cmd {
	switch castedComponent := component.(type) {
	case InteractiveComponent:
		return castedComponent.Update(msg)
	}
	return nil
}
