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
	"github.com/mieubrisse/teact/teact/utilities"
	"strconv"
)

type identificationFormImpl struct {
	components.Component

	nameInput text_input.TextInput
	ageInput  text_input.TextInput

	// TDOO extract this into something common, that all components can use??
	focusableItems []components.InteractiveComponent
	focusedItemIdx int

	isFocused bool
}

func New(opts ...IdentificationFormOpts) IdentificationForm {
	nameInput := text_input.New()
	ageInput := text_input.New()

	root := stylebox.New(
		flexbox.NewWithOpts(
			[]flexbox_item.FlexboxItem{
				flexbox_item.New(
					stylebox.New(
						text.New("IDENTIFICATION", text.WithAlign(text.AlignCenter)),
						stylebox.WithStyle(
							utilities.WithBold(true),
							utilities.WithForeground(colors.Platinum),
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
									stylebox.WithStyle(
										utilities.WithForeground(colors.Platinum),
										utilities.WithBold(true),
									),
								),
							),
							flexbox_item.New(nameInput),
						},
						flexbox.WithHorizontalAlignment(flexbox.AlignCenter),
					),
					flexbox_item.WithMinWidth(flexbox_item.FixedSize(10)),
				),
				flexbox_item.New(
					flexbox.NewWithOpts(
						[]flexbox_item.FlexboxItem{
							flexbox_item.New(
								stylebox.New(
									text.New("Age: "),
									stylebox.WithStyle(
										utilities.WithForeground(colors.Platinum),
										utilities.WithBold(true),
									),
								),
							),
							flexbox_item.New(ageInput),
						},
						flexbox.WithHorizontalAlignment(flexbox.AlignCenter),
					),
					flexbox_item.WithMinWidth(flexbox_item.FixedSize(10)),
				),
			},
			flexbox.WithDirection(flexbox.Column),
			flexbox.WithHorizontalAlignment(flexbox.AlignCenter),
		),
		stylebox.WithStyle(
			utilities.WithPadding(0, 1, 0, 1),
			utilities.WithBorder(lipgloss.NormalBorder()),
		),
	)

	result := &identificationFormImpl{
		Component: root,
		nameInput: nameInput,
		ageInput:  ageInput,
		focusableItems: []components.InteractiveComponent{
			nameInput,
			ageInput,
		},
		focusedItemIdx: 0,
		isFocused:      false,
	}
	for _, opt := range opts {
		opt(result)
	}
	return result
}

func (f identificationFormImpl) GetName() string {
	return f.nameInput.GetValue()
}

func (f *identificationFormImpl) SetName(name string) IdentificationForm {
	f.nameInput.SetValue(name)
	return f
}

func (f identificationFormImpl) GetAge() int {
	ageStr := f.ageInput.GetValue()
	result, _ := strconv.ParseInt(ageStr, 10, 64)
	return int(result)
}

func (f *identificationFormImpl) SetAge(age int) IdentificationForm {
	f.ageInput.SetValue(strconv.Itoa(age))
	return f
}

func (f *identificationFormImpl) Update(msg tea.Msg) tea.Cmd {
	if !f.isFocused {
		return nil
	}

	msgStr := utilities.GetMaybeKeyMsgStr(msg)
	switch msgStr {
	case "tab":
		// TODO extract this logic into something else

		newIdx := (f.focusedItemIdx + 1) % len(f.focusableItems)
		f.focusableItems[f.focusedItemIdx].SetFocus(false)
		f.focusableItems[newIdx].SetFocus(true)
		f.focusedItemIdx = newIdx
		return nil
	}

	return f.focusableItems[f.focusedItemIdx].Update(msg)
}

func (f *identificationFormImpl) SetFocus(isFocused bool) tea.Cmd {
	f.isFocused = isFocused
	return f.focusableItems[f.focusedItemIdx].SetFocus(isFocused)
}

func (f *identificationFormImpl) IsFocused() bool {
	return f.isFocused
}
