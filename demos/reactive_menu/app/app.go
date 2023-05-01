package app

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox"
	"github.com/mieubrisse/teact/teact/components/flexbox_item"
	"github.com/mieubrisse/teact/teact/components/list"
	"github.com/mieubrisse/teact/teact/components/stylebox"
	"github.com/mieubrisse/teact/teact/components/text"
	"github.com/mieubrisse/teact/teact/utilities"
)

// This is an app with a sidebar (like we'd see on websites), but which switches to a stacked orientation when
// the viewport is too small (like you'd see on mobile)
type ReactiveMenuApp interface {
	components.Component
}

type impl struct {
	// Root component
	components.Component

	box flexbox.Flexbox
}

func New() ReactiveMenuApp {
	menu := stylebox.New(
		list.NewWithContents[text.Text](
			text.New("Home", text.WithAlign(text.AlignCenter)),
			text.New("Search", text.WithAlign(text.AlignCenter)),
			text.New("Docs", text.WithAlign(text.AlignCenter)),
			text.New("About", text.WithAlign(text.AlignCenter)),
		).SetHorizontalAlignment(flexbox.AlignCenter),
		stylebox.WithExistingStyle(utilities.NewStyle(
			utilities.WithBorder(lipgloss.NormalBorder()),
			utilities.WithPadding(0, 1),
		)),
	)

	content := stylebox.New(
		text.New(
			"Four score and seven years ago our fathers brought forth "+
				"on this continent, a new nation, conceived in Liberty, and dedicated to the "+
				"proposition that all men are created equal.",
		),
		stylebox.WithExistingStyle(utilities.NewStyle(
			utilities.WithPadding(0, 1, 0, 1),
			utilities.WithBorder(lipgloss.NormalBorder()),
		)),
	)

	box := flexbox.New(
		flexbox_item.New(
			menu,
			flexbox_item.WithMaxWidth(flexbox_item.FixedSize(20)),
			flexbox_item.WithHorizontalGrowthFactor(2),
			flexbox_item.WithVerticalGrowthFactor(1),
		),
		flexbox_item.New(
			content,
			flexbox_item.WithHorizontalGrowthFactor(5),
			flexbox_item.WithVerticalGrowthFactor(1),
		),
	)

	return &impl{
		Component: box,
		box:       box,
	}
}

func (impl *impl) SetWidthAndGetDesiredHeight(actualWidth int) int {
	if actualWidth >= 60 {
		impl.box.SetDirection(flexbox.Row)
	} else {
		impl.box.SetDirection(flexbox.Column)
	}
	return impl.Component.SetWidthAndGetDesiredHeight(actualWidth)
}
