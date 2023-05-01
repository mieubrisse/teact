package identification_form

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/demos/reactive_form/colors"
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox"
	"github.com/mieubrisse/teact/teact/components/flexbox_item"
	"github.com/mieubrisse/teact/teact/components/stylebox"
	"github.com/mieubrisse/teact/teact/components/text"
	"github.com/mieubrisse/teact/teact/components/text_input"
	"github.com/mieubrisse/teact/teact/style"
)

type identificationFormImpl struct {
	components.Component

	input text_input.TextInput

	isFocused bool
}

func New(opts ...IdentificationFormOpts) IdentificationForm {
	input := text_input.New(text_input.WithFocus(true))

	root := stylebox.New(
		flexbox.NewWithOpts(
			[]flexbox_item.FlexboxItem{
				flexbox_item.New(
					stylebox.New(
						text.New("Identification", text.WithAlign(text.AlignCenter)),
						stylebox.WithNewStyle(
							style.WithBold(true),
							style.WithForeground(colors.Platinum),
						),
					),
					flexbox_item.WithHorizontalGrowthFactor(1),
				),
				flexbox_item.New(
					flexbox.NewWithOpts(
						[]flexbox_item.FlexboxItem{
							flexbox_item.New(
								stylebox.New(
									text.New("Name: "),
									stylebox.WithNewStyle(
										style.WithForeground(colors.Platinum),
									),
								),
							),
							flexbox_item.New(input),
						},
						flexbox.WithHorizontalAlignment(flexbox.AlignCenter),
					),
				),
			},
			flexbox.WithDirection(flexbox.Column),
			flexbox.WithHorizontalAlignment(flexbox.AlignCenter),
		),
		stylebox.WithNewStyle(
			style.WithBorder(lipgloss.NormalBorder()),
		),
	)

	result := &identificationFormImpl{
		Component: root,
		input:     input,
		isFocused: false,
	}
	for _, opt := range opts {
		opt(result)
	}
	return result
}

func (f identificationFormImpl) GetName() string {
	return f.input.GetValue()
}

func (f *identificationFormImpl) SetName(name string) IdentificationForm {
	f.input.SetValue(name)
	return f
}

func (f identificationFormImpl) Update(msg tea.Msg) tea.Cmd {
	if !f.isFocused {
		return nil
	}

	return f.input.Update(msg)
}

func (f *identificationFormImpl) SetFocus(isFocused bool) tea.Cmd {
	f.isFocused = isFocused
	if isFocused {

	}

	return f.input.SetFocus(isFocused)
}

func (f *identificationFormImpl) IsFocused() bool {
	return f.isFocused
}
