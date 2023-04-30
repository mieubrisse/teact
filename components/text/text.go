package text

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/box-layout-test/components"
	"github.com/muesli/reflow/ansi"
	"github.com/muesli/reflow/wordwrap"
	"strings"
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

type textImpl struct {
	text string

	alignment TextAlignment
}

func New(text string) Text {
	return &textImpl{
		text:      text,
		alignment: AlignLeft,
	}
}

func (t textImpl) GetContents() string {
	return t.text
}

func (t *textImpl) SetContents(str string) Text {
	t.text = str
	return t
}

func (t textImpl) GetTextAlignment() TextAlignment {
	return t.alignment
}

func (t *textImpl) SetTextAlignment(align TextAlignment) Text {
	t.alignment = align
	return t
}

func (t *textImpl) GetContentMinMax() (minWidth int, maxWidth int, minHeight int, maxHeight int) {
	minWidth = 0
	for _, field := range strings.Fields(t.text) {
		printableWidth := ansi.PrintableRuneWidth(field)
		if printableWidth > minWidth {
			minWidth = printableWidth
		}
	}

	maxWidth = lipgloss.Width(t.text)

	minHeight = lipgloss.Height(t.text)

	minWidthWrapped := wordwrap.String(t.text, minWidth)
	maxHeight = lipgloss.Height(minWidthWrapped)

	return
}

func (t textImpl) GetContentHeightForGivenWidth(width int) int {
	if width == 0 {
		return 0
	}

	// TODO cache this?
	wrapped := wordwrap.String(t.text, width)
	return lipgloss.Height(wrapped)
}

func (t textImpl) View(width int, height int) string {
	if width == 0 || height == 0 {
		return ""
	}

	wrapped := wordwrap.String(t.text, width)
	return lipgloss.NewStyle().Align(lipgloss.Position(t.alignment)).
		// Width to expand to a block
		Width(width).
		// Truncate (we can't support overrun or any other behaviours)
		MaxHeight(height).
		Render(wrapped)
}

// ====================================================================================================
//                                   Private Helper Functions
// ====================================================================================================
