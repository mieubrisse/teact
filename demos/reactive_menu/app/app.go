package app

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/components"
	flexbox2 "github.com/mieubrisse/teact/teact/components/flexbox"
	flexbox_item2 "github.com/mieubrisse/teact/teact/components/flexbox_item"
	"github.com/mieubrisse/teact/teact/components/list"
	"github.com/mieubrisse/teact/teact/components/stylebox"
	"github.com/mieubrisse/teact/teact/components/text"
	"github.com/mieubrisse/teact/teact/style"
)

// This is an app with a sidebar (like we'd see on websites), but which switches to a stacked orientation when
// the viewport is too small (like you'd see on mobile)
type ReactiveMenuApp interface {
	components.Component
}

type impl struct {
	// Root component
	components.Component

	box flexbox2.Flexbox
}

func New() ReactiveMenuApp {
	menu := stylebox.New(
		list.NewWithContents[text.Text](
			text.New("Home", text.WithAlign(text.AlignCenter)),
			text.New("Search", text.WithAlign(text.AlignCenter)),
			text.New("Docs", text.WithAlign(text.AlignCenter)),
			text.New("About", text.WithAlign(text.AlignCenter)),
		).SetHorizontalAlignment(flexbox2.AlignCenter),
		stylebox.WithExistingStyle(style.NewStyle(
			style.WithBorder(lipgloss.NormalBorder()),
			style.WithPadding(0, 1),
		)),
	)

	content := stylebox.New(
		text.New(
			"Four score and seven years ago our fathers brought forth "+
				"on this continent, a new nation, conceived in Liberty, and dedicated to the "+
				"proposition that all men are created equal.",
		),
		stylebox.WithExistingStyle(style.NewStyle(
			style.WithPadding(0, 1, 0, 1),
			style.WithBorder(lipgloss.NormalBorder()),
		)),
	)

	box := flexbox2.NewWithContents(
		flexbox_item2.New(
			menu,
			flexbox_item2.WithMaxWidth(flexbox_item2.FixedSize(20)),
			flexbox_item2.WithHorizontalGrowthFactor(2),
			flexbox_item2.WithVerticalGrowthFactor(1),
		),
		flexbox_item2.New(
			content,
			flexbox_item2.WithHorizontalGrowthFactor(5),
			flexbox_item2.WithVerticalGrowthFactor(1),
		),
	)

	return &impl{
		Component: box,
		box:       box,
	}
}

func (impl *impl) SetWidthAndGetDesiredHeight(actualWidth int) int {
	if actualWidth >= 60 {
		impl.box.SetDirection(flexbox2.Row)
	} else {
		impl.box.SetDirection(flexbox2.Column)
	}
	return impl.Component.SetWidthAndGetDesiredHeight(actualWidth)
}
