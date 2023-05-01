package identification_form

import "github.com/mieubrisse/teact/teact/components"

// TODO build a form component in Teact itself once we have Grid layout
type IdentificationForm interface {
	components.InteractiveComponent

	GetName() string
	SetName(name string) IdentificationForm

	GetAge() int
	SetAge(age int) IdentificationForm
}
