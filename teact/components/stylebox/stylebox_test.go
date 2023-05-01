package stylebox

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/component_test"
	"github.com/mieubrisse/teact/teact/components/text"
	"testing"
)

var inner = text.New("\nThis is a\nmultiline string\n\n")
var innerMinWidth, innerMaxWidth, innerMinHeight, innerMaxHeight = inner.GetContentMinMax()
var noChangeAssertion = component_test.GetRenderedContentAssertion(
	innerMaxWidth,
	innerMinHeight,
	inner.View(innerMaxWidth, innerMaxHeight),
)

func TestUnstyled(t *testing.T) {
	component := New(inner)

	assertions := component_test.FlattenAssertionGroups(
		component_test.GetDefaultAssertions(),
		component_test.GetContentSizeAssertions(innerMinWidth, innerMaxWidth, innerMinHeight, innerMaxHeight),
		noChangeAssertion,
	)
	component_test.CheckAll(
		t,
		assertions,
		component,
	)
}

func TestPadding(t *testing.T) {
	// Even padding
	padding := 2
	component := New(inner).SetStyle(lipgloss.NewStyle().Padding(padding))

	assertions := component_test.FlattenAssertionGroups(
		component_test.GetContentSizeAssertions(
			2*padding+innerMinWidth,
			2*padding+innerMaxWidth,
			2*padding+innerMinHeight,
			2*padding+innerMaxHeight,
		),
		// Should be only padding when there's no place for content
		component_test.GetRenderedContentAssertion(3, 3, "   \n   \n   "),
		component_test.GetRenderedContentAssertion(5, 6, "     \n     \n     \n  T  \n     \n     "),
	)

	component_test.CheckAll(
		t,
		assertions,
		component,
	)
}

func TestBorder(t *testing.T) {
	style := lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())
	component := New(inner).SetStyle(style)

	assertions := component_test.FlattenAssertionGroups(
		component_test.GetContentSizeAssertions(
			innerMinWidth+style.GetHorizontalBorderSize(),
			innerMaxWidth+style.GetHorizontalBorderSize(),
			innerMinHeight+style.GetVerticalBorderSize(),
			innerMaxHeight+style.GetVerticalBorderSize(),
		),
		component_test.GetHeightAtWidthAssertions(
			innerMaxWidth+style.GetVerticalBorderSize(),
			innerMinHeight+style.GetVerticalBorderSize(),
		),
		component_test.GetHeightAtWidthAssertions(
			innerMinWidth+style.GetVerticalBorderSize(),
			innerMaxHeight+style.GetVerticalBorderSize(),
		),
	)

	component_test.CheckAll(
		t,
		assertions,
		component,
	)
}

func TestColorStylesMaintainSize(t *testing.T) {
	styles := []lipgloss.Style{
		lipgloss.NewStyle(),
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF22FF")).
			Background(lipgloss.Color("#333333")).
			Bold(true).
			Faint(true).
			Blink(true).
			UnderlineSpaces(true).
			Underline(true).
			Italic(true).
			ColorWhitespace(true),
	}

	for _, style := range styles {
		component := New(inner).SetStyle(style)
		component_test.CheckAll(t, noChangeAssertion, component)
	}
}

func TestProhibitedStylesAreRemoved(t *testing.T) {
	prohibitedStyles := []lipgloss.Style{
		lipgloss.NewStyle().Margin(2),
		lipgloss.NewStyle().Align(lipgloss.Center),
		lipgloss.NewStyle().AlignHorizontal(lipgloss.Center),
		lipgloss.NewStyle().AlignVertical(lipgloss.Center),
		lipgloss.NewStyle().Width(1),
		lipgloss.NewStyle().MaxWidth(1),
		lipgloss.NewStyle().Height(1),
		lipgloss.NewStyle().MaxHeight(1),
		lipgloss.NewStyle().Inline(true),
	}

	for _, style := range prohibitedStyles {
		component := New(inner).SetStyle(style)
		component_test.CheckAll(t, noChangeAssertion, component)
	}
}
