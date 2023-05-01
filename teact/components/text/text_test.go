package text

import (
	"github.com/mieubrisse/teact/teact/component_test"
	"testing"
)

func TestShortString(t *testing.T) {
	str := "This is a short string"
	minWidth := 6
	maxWidth := 22
	minHeight := 1
	maxHeight := 4

	sizeAssertions := component_test.FlattenAssertionGroups(
		component_test.GetDefaultAssertions(),
		component_test.GetContentSizeAssertions(minWidth, maxWidth, minHeight, maxHeight),
		component_test.GetHeightAtWidthAssertions(
			minWidth, maxHeight, // min content width
			8, 3, // in the middle
			maxWidth, minHeight, // max content width
			100, minHeight, // beyond max content width
		),
	)

	// Verify that the size assertions are valid at all alignments
	for _, alignment := range []TextAlignment{AlignLeft, AlignCenter, AlignRight} {
		component := New(str).SetTextAlignment(alignment)
		component_test.CheckAll(t, sizeAssertions, component)
	}
}

func TestStringWithNewlines(t *testing.T) {
	str := "This is the first line\nHere's a second\nAnd a third"
	minWidth := 6
	maxWidth := 22
	minHeight := 3
	maxHeight := 9

	sizeAssertions := component_test.FlattenAssertionGroups(
		component_test.GetDefaultAssertions(),
		component_test.GetContentSizeAssertions(minWidth, maxWidth, minHeight, maxHeight),
		component_test.GetHeightAtWidthAssertions(
			0, 0, // invisible
			minWidth, maxHeight, // min content width
			10, 7, // in the middle
			maxWidth, minHeight, // max content width
			100, minHeight, // beyond max content width
		),
	)

	// Verify that the size assertions are valid at all alignments
	for _, alignment := range []TextAlignment{AlignLeft, AlignCenter, AlignRight} {
		component := New(str).SetTextAlignment(alignment)
		component_test.CheckAll(t, sizeAssertions, component)
	}
}

func TestInvisibleString(t *testing.T) {
	str := ""

	sizeAssertions := component_test.FlattenAssertionGroups(
		component_test.GetDefaultAssertions(),
		component_test.GetContentSizeAssertions(0, 0, 1, 1),
		component_test.GetHeightAtWidthAssertions(
			0, 1,
			1, 1,
			10, 1,
		),
	)

	component_test.CheckAll(t, sizeAssertions, New(str))
}

func TestSmallWidths(t *testing.T) {
	text := "\nThis is a\nmultiline string\n\n"
	component := New(text)

	assertions := component_test.FlattenAssertionGroups(
		component_test.GetRenderedContentAssertion(2, 2, "  \nTh"),
		component_test.GetRenderedContentAssertion(2, 5, "  \nTh\nis\nis\na "),
		component_test.GetRenderedContentAssertion(5, 5, "     \nThis \nis a \nmulti\nline "),
	)

	component_test.CheckAll(t, assertions, component)
}
