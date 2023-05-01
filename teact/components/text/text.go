package text

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/components"
)

type TextAlignment lipgloss.Position

const (
	AlignLeft   = TextAlignment(lipgloss.Left)
	AlignCenter = TextAlignment(lipgloss.Center)
	AlignRight  = TextAlignment(lipgloss.Right)
)

// Analogous to the <p> tag in HTML
type Text interface {
	components.Component

	GetContents() string
	SetContents(str string) Text

	GetTextAlignment() TextAlignment
	SetTextAlignment(alignment TextAlignment) Text
}
