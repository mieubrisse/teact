package text

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/utilities"
	"github.com/muesli/reflow/ansi"
	"github.com/muesli/reflow/wordwrap"
	"strings"
)

type textImpl struct {
	contents string

	alignment TextAlignment
}

func New(opts ...TextOpt) Text {
	result := &textImpl{
		contents:  "",
		alignment: AlignLeft,
	}
	for _, opt := range opts {
		opt(result)
	}
	return result
}

func (t textImpl) GetContents() string {
	return t.contents
}

func (t *textImpl) SetContents(str string) Text {
	t.contents = str
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
	for _, field := range strings.Fields(t.contents) {
		printableWidth := ansi.PrintableRuneWidth(field)
		if printableWidth > minWidth {
			minWidth = printableWidth
		}
	}

	maxWidth = lipgloss.Width(t.contents)

	maxHeight = t.SetWidthAndGetDesiredHeight(minWidth)
	minHeight = t.SetWidthAndGetDesiredHeight(maxWidth)

	return
}

func (t textImpl) SetWidthAndGetDesiredHeight(width int) int {
	if width == 0 {
		return 0
	}

	// TODO cache this?
	wrapped := wordwrap.String(t.contents, width)
	return lipgloss.Height(wrapped)
}

func (t textImpl) View(width int, height int) string {
	if width == 0 || height == 0 {
		return ""
	}

	result := wordwrap.String(t.contents, width)

	// Ensure we have a string no more than max (though it may still be short)
	result = lipgloss.NewStyle().MaxWidth(width).MaxHeight(height).Render(result)

	// Now align (the string may still be short)
	result = lipgloss.NewStyle().Align(lipgloss.Position(t.alignment)).Render(result)

	// Place in the correct location
	result = lipgloss.PlaceHorizontal(width, lipgloss.Position(t.alignment), result)

	// Finally, coerce to expand if necessary
	result = utilities.Coerce(result, width, height)

	return result
}

// ====================================================================================================
//                                   Private Helper Functions
// ====================================================================================================
