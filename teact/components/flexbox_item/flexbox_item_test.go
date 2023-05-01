package flexbox_item

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/components/stylebox"
	test_assertions2 "github.com/mieubrisse/teact/teact/components/test_assertions"
	"github.com/mieubrisse/teact/teact/components/text"
	"testing"
)

var inner = text.New("\nThis is a\nmultiline string\n\n")
var innerMinWidth, innerMaxWidth, innerMinHeight, innerMaxHeight = inner.GetContentMinMax()
var noChangeAssertion = test_assertions2.RenderedContentAssertion{
	Width:           innerMaxWidth,
	Height:          innerMinHeight,
	ExpectedContent: inner.View(innerMaxWidth, innerMaxHeight),
}

func TestBasic(t *testing.T) {
	component := New(inner)
	assertions := test_assertions2.FlattenAssertionGroups(
		test_assertions2.GetDefaultAssertions(),
		test_assertions2.GetContentSizeAssertions(innerMinWidth, innerMaxWidth, innerMinHeight, innerMaxHeight),
		test_assertions2.GetHeightAtWidthAssertions(
			1, inner.SetWidthAndGetDesiredHeight(1),
			2, inner.SetWidthAndGetDesiredHeight(2),
			10, inner.SetWidthAndGetDesiredHeight(10),
		),
		[]test_assertions2.ComponentAssertion{noChangeAssertion},
	)

	test_assertions2.CheckAll(t, assertions, component)
}

func TestTruncate(t *testing.T) {
	component := New(inner).SetOverflowStyle(Truncate)
	assertions := test_assertions2.FlattenAssertionGroups(
		test_assertions2.GetRenderedContentAssertion(4, 4, "    \nThis\nmult\n    "),
	)

	test_assertions2.CheckAll(t, assertions, component)
}

func TestWrap(t *testing.T) {
	component := New(inner).SetOverflowStyle(Wrap)
	assertions := test_assertions2.FlattenAssertionGroups(
		test_assertions2.GetRenderedContentAssertion(4, 4, "    \nThis\nis a\nmult"),
	)

	test_assertions2.CheckAll(t, assertions, component)
}

func TestStyleboxInside(t *testing.T) {
	contained := stylebox.New(text.New("This is child 2")).
		SetStyle(lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()))
	component := New(contained)

	assertions := test_assertions2.FlattenAssertionGroups(
		test_assertions2.GetContentSizeAssertions(7, 17, 3, 6),
	)

	test_assertions2.CheckAll(t, assertions, component)

}

func TestFixedSize(t *testing.T) {
	fixedWidth := 60

	contained := text.New("A description of pizza")
	item := New(contained).
		SetMinWidth(FixedSize(fixedWidth)).
		SetMaxWidth(FixedSize(fixedWidth))

	assertions := test_assertions2.FlattenAssertionGroups(
		test_assertions2.GetContentSizeAssertions(fixedWidth, fixedWidth, 1, 1),
	)

	test_assertions2.CheckAll(t, assertions, item)
}
