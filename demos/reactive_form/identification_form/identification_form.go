package identification_form

import "github.com/mieubrisse/teact/teact/components"

type IdentificationForm interface {
	components.InteractiveComponent

	GetName() string
	SetName(name string) IdentificationForm

	GetAge() int
	SetAge(age int) IdentificationForm
}
