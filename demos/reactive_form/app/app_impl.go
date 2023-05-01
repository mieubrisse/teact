package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox"
	"github.com/mieubrisse/teact/teact/components/flexbox_item"
	"github.com/mieubrisse/teact/teact/components/stylebox"
	"github.com/mieubrisse/teact/teact/components/text"
	"github.com/mieubrisse/teact/teact/components/text_input"
	"github.com/mieubrisse/teact/teact/style"
)

type reactiveFormAppImpl struct {
	components.Component

	input text_input.TextInput
}

func New() ReactiveFormApp {
	input := text_input.New(text_input.WithFocus(true))
	app := flexbox.NewWithOpts(
		[]flexbox_item.FlexboxItem{
			flexbox_item.New(
				stylebox.New(
					flexbox.NewWithOpts(
						[]flexbox_item.FlexboxItem{
							flexbox_item.New(text.New("Form")),
							flexbox_item.New(
								flexbox.NewWithOpts(
									[]flexbox_item.FlexboxItem{
										flexbox_item.New(text.New("Name: ")),
										flexbox_item.New(input),
									},
									flexbox.WithHorizontalAlignment(flexbox.AlignCenter),
								),
							),
						},
						flexbox.WithDirection(flexbox.Column),
						flexbox.WithHorizontalAlignment(flexbox.AlignCenter),
					),
					stylebox.WithStyle(
						style.NewStyle(style.WithBorder(lipgloss.NormalBorder())),
					),
				),
			),
		},
		flexbox.WithDirection(flexbox.Column),
	)

	return &reactiveFormAppImpl{
		Component: app,
	}
}

func (r reactiveFormAppImpl) Update(msg tea.Msg) tea.Cmd {
	return r.input.Update(msg)
}

func (r reactiveFormAppImpl) SetFocus(isFocused bool) tea.Cmd {
	return nil
}

func (r reactiveFormAppImpl) IsFocused() bool {
	return true
}
