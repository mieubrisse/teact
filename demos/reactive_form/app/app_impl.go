package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/demos/reactive_form/bio_card"
	"github.com/mieubrisse/teact/demos/reactive_form/identification_form"
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox"
	"github.com/mieubrisse/teact/teact/components/flexbox_item"
)

type reactiveFormAppImpl struct {
	components.Component

	form    identification_form.IdentificationForm
	bioCard bio_card.BioCard

	box            flexbox.Flexbox
	formBoxItem    flexbox_item.FlexboxItem
	bioCardBoxItem flexbox_item.FlexboxItem
}

func New() ReactiveFormApp {
	form := identification_form.New(
		identification_form.WithFocus(true),
		identification_form.WithName("007"),
	)
	formBoxItem := flexbox_item.New(
		form,
		flexbox_item.WithHorizontalGrowthFactor(2),
		flexbox_item.WithVerticalGrowthFactor(1),
	)

	bioCard := bio_card.New()
	bioCardBoxItem := flexbox_item.New(
		bioCard,
		flexbox_item.WithHorizontalGrowthFactor(5),
		flexbox_item.WithVerticalGrowthFactor(1),
	)
	box := flexbox.New(formBoxItem, bioCardBoxItem)

	result := &reactiveFormAppImpl{
		Component:      box,
		form:           form,
		bioCard:        bioCard,
		box:            box,
		formBoxItem:    formBoxItem,
		bioCardBoxItem: bioCardBoxItem,
	}
	result.updateBioCard()
	return result
}

func (r *reactiveFormAppImpl) SetWidthAndGetDesiredHeight(actualWidth int) int {
	if actualWidth < 80 {
		r.box.SetDirection(flexbox.Column)
		r.formBoxItem.SetVerticalGrowthFactor(0)
	} else {
		r.box.SetDirection(flexbox.Row)
		r.formBoxItem.SetVerticalGrowthFactor(1)
	}
	return r.Component.SetWidthAndGetDesiredHeight(actualWidth)
}

func (r reactiveFormAppImpl) Update(msg tea.Msg) tea.Cmd {
	result := r.form.Update(msg)
	r.updateBioCard()
	return result
}

func (r reactiveFormAppImpl) SetFocus(isFocused bool) tea.Cmd {
	return nil
}

func (r reactiveFormAppImpl) IsFocused() bool {
	return true
}

// ====================================================================================================
//
//	Private Helper Functions
//
// ====================================================================================================
func (r *reactiveFormAppImpl) updateBioCard() {
	r.bioCard.SetName(r.form.GetName())
}
