package text_input

import "github.com/mieubrisse/teact/teact/components"

type TextInput interface {
	components.InteractiveComponent

	GetValue() string
	SetValue(value string) TextInput
}
