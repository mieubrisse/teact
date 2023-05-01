package input

import "github.com/mieubrisse/teact/teact/components"

type Input interface {
	components.InteractiveComponent

	GetValue() string
	SetValue(value string) Input
}
