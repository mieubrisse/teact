package stylebox

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/utilities"
)

type StyleboxOpt func(Stylebox)

// Convenience function for styling the stylebox with a new lipgloss.Style
func WithStyle(styleOpts ...utilities.StyleOpt) StyleboxOpt {
	return func(box Stylebox) {
		newStyle := lipgloss.NewStyle()
		for _, opt := range styleOpts {
			newStyle = opt(newStyle)
		}
		box.SetStyle(newStyle)
	}
}

// Set the stylebox's style to an existing lipgloss.Style
func WithExistingStyle(newStyle lipgloss.Style) StyleboxOpt {
	return func(box Stylebox) {
		box.SetStyle(newStyle)
	}
}
