package stylebox

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/style"
)

type StyleboxOpt func(Stylebox)

func WithNewStyle(styleOpts ...style.StyleOpt) StyleboxOpt {
	return func(box Stylebox) {
		newStyle := lipgloss.NewStyle()
		for _, opt := range styleOpts {
			newStyle = opt(newStyle)
		}
		box.SetStyle(newStyle)
	}
}

func WithStyle(newStyle lipgloss.Style) StyleboxOpt {
	return func(box Stylebox) {
		box.SetStyle(newStyle)
	}
}
