package stylebox

import "github.com/charmbracelet/lipgloss"

type StyleboxOpt func(stylebox Stylebox)

func WithStyle(style lipgloss.Style) StyleboxOpt {
	return func(box Stylebox) {
		box.SetStyle(style)
	}
}
