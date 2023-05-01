package stylebox

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/components"
)

// Stylebox is a box explicitly for controlling an element's style
// No other elements control style; this is intentional so that
// it's clear where exactly style is getting changed
type Stylebox interface {
	components.Component

	GetStyle() lipgloss.Style
	// NOTE: all layout-affecting properties (height, width, alignment, margin, inline) are ignored
	// The only layout-affecting property left in place are border and padding
	SetStyle(style lipgloss.Style) Stylebox
}
