package flexbox_item

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/component_test"
	"github.com/mieubrisse/teact/teact/components/stylebox"
	"github.com/mieubrisse/teact/teact/components/text"
	"testing"
)

var inner = text.New("\nThis is a\nmultiline string\n\n")
var innerMinWidth, innerMaxWidth, innerMinHeight, innerMaxHeight = inner.GetContentMinMax()
var noChangeAssertion = component_test.RenderedContentAssertion{
	Width:           innerMaxWidth,
	Height:          innerMinHeight,
	ExpectedContent: inner.View(innerMaxWidth, innerMaxHeight),
}

func TestBasic(t *testing.T) {
	component := New(inner)
	assertions := component_test.FlattenAssertionGroups(
		component_test.GetDefaultAssertions(),
		component_test.GetContentSizeAssertions(innerMinWidth, innerMaxWidth, innerMinHeight, innerMaxHeight),
		component_test.GetHeightAtWidthAssertions(
			1, inner.SetWidthAndGetDesiredHeight(1),
			2, inner.SetWidthAndGetDesiredHeight(2),
			10, inner.SetWidthAndGetDesiredHeight(10),
		),
		[]component_test.ComponentAssertion{noChangeAssertion},
	)

	component_test.CheckAll(t, assertions, component)
}

func TestTruncate(t *testing.T) {
	component := New(inner).SetOverflowStyle(Truncate)
	assertions := component_test.FlattenAssertionGroups(
		component_test.GetRenderedContentAssertion(4, 4, "    \nThis\nmult\n    "),
	)

	component_test.CheckAll(t, assertions, component)
}

func TestWrap(t *testing.T) {
	component := New(inner).SetOverflowStyle(Wrap)
	assertions := component_test.FlattenAssertionGroups(
		component_test.GetRenderedContentAssertion(4, 4, "    \nThis\nis a\nmult"),
	)

	component_test.CheckAll(t, assertions, component)
}

func TestStyleboxInside(t *testing.T) {
	contained := stylebox.New(text.New("This is child 2")).
		SetStyle(lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()))
	component := New(contained)

	assertions := component_test.FlattenAssertionGroups(
		component_test.GetContentSizeAssertions(7, 17, 3, 6),
	)

	component_test.CheckAll(t, assertions, component)

}

func TestFixedSize(t *testing.T) {
	fixedWidth := 60

	contained := text.New("A description of pizza")
	item := New(contained).
		SetMinWidth(FixedSize(fixedWidth)).
		SetMaxWidth(FixedSize(fixedWidth))

	assertions := component_test.FlattenAssertionGroups(
		component_test.GetContentSizeAssertions(fixedWidth, fixedWidth, 1, 1),
	)

	component_test.CheckAll(t, assertions, item)
}
