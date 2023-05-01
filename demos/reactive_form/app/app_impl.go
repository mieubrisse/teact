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
}

func New() ReactiveFormApp {
	form := identification_form.New(
		identification_form.WithFocus(true),
		identification_form.WithName("007"),
	)
	bioCard := bio_card.New()
	app := flexbox.NewWithOpts(
		[]flexbox_item.FlexboxItem{
			flexbox_item.New(
				form,
				flexbox_item.WithHorizontalGrowthFactor(1),
				flexbox_item.WithVerticalGrowthFactor(1),
			),
			flexbox_item.New(
				bioCard,
				flexbox_item.WithHorizontalGrowthFactor(1),
				flexbox_item.WithVerticalGrowthFactor(1),
			),
		},
		flexbox.WithDirection(flexbox.Column),
	)

	result := &reactiveFormAppImpl{
		Component: app,
		form:      form,
		bioCard:   bioCard,
	}
	result.updateBioCard()
	return result
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
