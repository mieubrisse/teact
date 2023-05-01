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

type HelloWorldApp interface {
	components.Component
}

type helloWorldAppImpl struct {
	// So long as we assign a component to this then our component will call down to it (via Go struct embedding)
	components.Component
}

func New() HelloWorldApp {
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
