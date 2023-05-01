package secret_agent_terminal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/demos/reactive_form/bio_card"
	"github.com/mieubrisse/teact/demos/reactive_form/colors"
	"github.com/mieubrisse/teact/demos/reactive_form/identification_form"
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox"
	"github.com/mieubrisse/teact/teact/components/flexbox_item"
	"github.com/mieubrisse/teact/teact/components/stylebox"
	"github.com/mieubrisse/teact/teact/components/text"
	"github.com/mieubrisse/teact/teact/utilities"
)

const (
	// Below this size, the form & bio card will be stacked vertically
	columnSwitchThreshold = 80
)

type secretAgentTerminalImpl struct {
	components.Component

	form    identification_form.IdentificationForm
	bioCard bio_card.BioCard

	contentBox     flexbox.Flexbox
	formBoxItem    flexbox_item.FlexboxItem
	bioCardBoxItem flexbox_item.FlexboxItem
}

func New() SecretAgentTerminal {
	form := identification_form.New(
		identification_form.WithFocus(true),
		identification_form.WithName("007"),
		identification_form.WithAge(55),
	)
	formBoxItem := flexbox_item.New(
		form,
		flexbox_item.WithMaxWidth(flexbox_item.FixedSize(40)),
		// growth factors will be handled upon render, based on viewport size
	)

	bioCard := bio_card.New()
	bioCardBoxItem := flexbox_item.New(
		bioCard,
		flexbox_item.WithHorizontalGrowthFactor(1),
		flexbox_item.WithVerticalGrowthFactor(1),
	)
	contentBox := flexbox.New(formBoxItem, bioCardBoxItem)

	appTitle := stylebox.New(
		text.New("SECRET AGENT TERMINAL APP", text.WithAlign(text.AlignCenter)),
		stylebox.WithStyle(
			utilities.WithForeground(colors.VividSkyBlue),
			utilities.WithBold(true),
		),
	)

	var root components.Component = flexbox.NewWithOpts(
		[]flexbox_item.FlexboxItem{
			flexbox_item.New(
				appTitle,
				flexbox_item.WithHorizontalGrowthFactor(1),
			),
			flexbox_item.New(
				contentBox,
				flexbox_item.WithHorizontalGrowthFactor(1),
				flexbox_item.WithVerticalGrowthFactor(1),
			),
		},
		flexbox.WithDirection(flexbox.Column),
	)

	root = stylebox.New(
		root,
		stylebox.WithStyle(utilities.WithPadding(1, 2, 1, 2)),
	)

	result := &secretAgentTerminalImpl{
		Component:      root,
		form:           form,
		bioCard:        bioCard,
		contentBox:     contentBox,
		formBoxItem:    formBoxItem,
		bioCardBoxItem: bioCardBoxItem,
	}
	result.updateBioCard()
	return result
}

func (terminal *secretAgentTerminalImpl) SetWidthAndGetDesiredHeight(actualWidth int) int {
	if actualWidth < columnSwitchThreshold {
		terminal.contentBox.SetDirection(flexbox.Column)
		terminal.formBoxItem.SetVerticalGrowthFactor(0)
		terminal.formBoxItem.SetHorizontalGrowthFactor(1)
	} else {
		terminal.contentBox.SetDirection(flexbox.Row)
		terminal.formBoxItem.SetVerticalGrowthFactor(1)
		terminal.formBoxItem.SetHorizontalGrowthFactor(0)
	}
	return terminal.Component.SetWidthAndGetDesiredHeight(actualWidth)
}

func (terminal secretAgentTerminalImpl) Update(msg tea.Msg) tea.Cmd {
	result := terminal.form.Update(msg)
	terminal.updateBioCard()
	return result
}

func (terminal secretAgentTerminalImpl) SetFocus(isFocused bool) tea.Cmd {
	return nil
}

func (terminal secretAgentTerminalImpl) IsFocused() bool {
	return true
}

// ====================================================================================================
//
//	Private Helper Functions
//
// ====================================================================================================
func (terminal *secretAgentTerminalImpl) updateBioCard() {
	terminal.bioCard.SetName(terminal.form.GetName())
	terminal.bioCard.SetAge(terminal.form.GetAge())
}
