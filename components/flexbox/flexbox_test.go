package flexbox

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/components/flexbox_item"
	"github.com/mieubrisse/teact/components/stylebox"
	"github.com/mieubrisse/teact/components/test_assertions"
	"github.com/mieubrisse/teact/components/text"
	"testing"
)

func TestColumnLayout(t *testing.T) {
	child1 := text.New("This is child 1")
	child2 := stylebox.New(text.New("This is child 2")).
		SetStyle(lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()))
	child3 := text.New("This is child 3")

	flexbox := NewWithContents(
		flexbox_item.New(child1),
		flexbox_item.New(child2),
		flexbox_item.New(child3),
	).SetHorizontalAlignment(AlignCenter).SetVerticalAlignment(AlignCenter).SetDirection(Column)

	width, height := 30, 30

	assertions := test_assertions.FlattenAssertionGroups(
		test_assertions.GetDefaultAssertions(),
		test_assertions.GetContentSizeAssertions(
			7,
			17,
			5,
			20,
		),
	)

	// Need to populate the caches
	flexbox.GetContentMinMax()
	flexbox.GetContentHeightForGivenWidth(width)
	flexbox.View(width, height)

	test_assertions.CheckAll(t, assertions, flexbox)
}

/*
func TestAdvancedColumnLayout(t *testing.T) {
	var red = lipgloss.Color("#FF0000")
	var blue = lipgloss.Color("#0000FF")
	var green = lipgloss.Color("#00FF00")
	var lightGray = lipgloss.Color("#333333")
	var text1Style = lipgloss.NewStyle().Foreground(red).Background(lightGray)
	var text2Style = lipgloss.NewStyle().Foreground(green).Border(lipgloss.NormalBorder())
	var text3Style = lipgloss.NewStyle().Foreground(blue).Background(lightGray)

	text1 := stylebox.New(text.New("This is text 1")).SetStyle(text1Style)
	text2 := stylebox.New(text.New("This is text 2")).SetStyle(text2Style)
	text3 := stylebox.New(
		text.New("Four score and seven years ago our fathers brought forth on this continent, " +
			"a new nation, conceived in Liberty, and dedicated to the proposition that all men " +
			"are created equal.").
			SetTextAlignment(text.AlignCenter)).SetStyle(text3Style)

	component := NewWithContents(
		flexbox_item.New(text1),
		flexbox_item.New(text2),
		flexbox_item.New(text3),
	).SetHorizontalAlignment(AlignCenter).
		SetVerticalAlignment(AlignCenter).SetDirection(Column)

	width, height := 30, 30
	component.GetContentMinMax()
	component.GetContentHeightForGivenWidth(width)
	component.View(width, height)
}
*/

func TestFixedSizeItem(t *testing.T) {
	nameText := text.New("Pizza")
	descriptionText := text.New("A description of pizza")

	style := lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())

	box := NewWithContents(
		flexbox_item.New(stylebox.New(nameText).SetStyle(style)).
			SetMinWidth(flexbox_item.FixedSize(60)).
			SetMaxWidth(flexbox_item.FixedSize(60)),
		flexbox_item.New(text.New(" ")),
		flexbox_item.New(stylebox.New(descriptionText).SetStyle(style)).SetMaxWidth(flexbox_item.MaxAvailable),
	)

	width, height := 170, 30

	box.GetContentMinMax()
	box.GetContentHeightForGivenWidth(width)
	box.View(width, height)
}
