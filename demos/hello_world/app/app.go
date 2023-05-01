package app

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox"
	"github.com/mieubrisse/teact/teact/components/flexbox_item"
	"github.com/mieubrisse/teact/teact/components/stylebox"
	"github.com/mieubrisse/teact/teact/components/text"
	"github.com/mieubrisse/teact/teact/style"
)

// A custom component
type HelloWorldApp interface {
	components.Component
}

// Implementation of the custom component
type helloWorldAppImpl struct {
	// So long as we assign a component to this then our component will call down to it (via Go struct embedding)
	components.Component
}

func New() HelloWorldApp {
	// This is a tree, just like HTML, with leaf nodes indented the most
	root := flexbox.NewWithOpts(
		[]flexbox_item.FlexboxItem{
			flexbox_item.New(
				stylebox.New(
					text.New("Hello, world!"),
					stylebox.WithStyle(
						style.WithForeground(lipgloss.Color("#B6DCFE")),
					),
				),
			),
		},
		flexbox.WithVerticalAlignment(flexbox.AlignCenter),
		flexbox.WithHorizontalAlignment(flexbox.AlignCenter),
	)

	return &helloWorldAppImpl{
		Component: root,
	}
}
