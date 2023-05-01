package text

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/ansi"
	"github.com/muesli/reflow/wordwrap"
	"strings"
)

type textImpl struct {
	text string

	alignment TextAlignment
}

func New(text string, opts ...TextOpt) Text {
	result := &textImpl{
		text:      text,
		alignment: AlignLeft,
	}
	for _, opt := range opts {
		opt(result)
	}
	return result
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

	maxHeight = t.SetWidthAndGetDesiredHeight(minWidth)
	minHeight = t.SetWidthAndGetDesiredHeight(maxWidth)

	return
}

func (t textImpl) SetWidthAndGetDesiredHeight(width int) int {
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
