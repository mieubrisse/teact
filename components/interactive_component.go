package components

import tea "github.com/charmbracelet/bubbletea"

type InteractiveComponent interface {
	Component

	// Update updates the model based on the given message
	// We do this by-reference, because by-value is just too messy (see README of this repo)
	// BLUF: you get into weird situations with generic interfaces
	Update(msg tea.Msg) tea.Cmd

	SetFocus(isFocused bool) tea.Cmd
	IsFocused() bool
}
