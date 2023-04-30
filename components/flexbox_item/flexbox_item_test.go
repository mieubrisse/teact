package flexbox_item

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/box-layout-test/components/stylebox"
	"github.com/mieubrisse/box-layout-test/components/test_assertions"
	"github.com/mieubrisse/box-layout-test/components/text"
	"testing"
)

var inner = text.New("\nThis is a\nmultiline string\n\n")
var innerMinWidth, innerMaxWidth, innerMinHeight, innerMaxHeight = inner.GetContentMinMax()
var noChangeAssertion = test_assertions.RenderedContentAssertion{
	Width:           innerMaxWidth,
	Height:          innerMinHeight,
	ExpectedContent: inner.View(innerMaxWidth, innerMaxHeight),
}

func TestBasic(t *testing.T) {
	component := New(inner)
	assertions := test_assertions.FlattenAssertionGroups(
		test_assertions.GetDefaultAssertions(),
		test_assertions.GetContentSizeAssertions(innerMinWidth, innerMaxWidth, innerMinHeight, innerMaxHeight),
		test_assertions.GetHeightAtWidthAssertions(
			1, inner.GetContentHeightForGivenWidth(1),
			2, inner.GetContentHeightForGivenWidth(2),
			10, inner.GetContentHeightForGivenWidth(10),
		),
		[]test_assertions.ComponentAssertion{noChangeAssertion},
	)

	test_assertions.CheckAll(t, assertions, component)
}

func TestTruncate(t *testing.T) {
	component := New(inner).SetOverflowStyle(Truncate)
	assertions := test_assertions.FlattenAssertionGroups(
		test_assertions.GetRenderedContentAssertion(4, 4, "    \nThis\nmult\n    "),
	)

	test_assertions.CheckAll(t, assertions, component)
}

func TestWrap(t *testing.T) {
	component := New(inner).SetOverflowStyle(Wrap)
	assertions := test_assertions.FlattenAssertionGroups(
		test_assertions.GetRenderedContentAssertion(4, 4, "    \nThis\nis a\nmult"),
	)

	test_assertions.CheckAll(t, assertions, component)
}

func TestStyleboxInside(t *testing.T) {
	contained := stylebox.New(text.New("This is child 2")).
		SetStyle(lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()))
	component := New(contained)

	assertions := test_assertions.FlattenAssertionGroups(
		test_assertions.GetContentSizeAssertions(7, 17, 3, 6),
	)

	test_assertions.CheckAll(t, assertions, component)

}
