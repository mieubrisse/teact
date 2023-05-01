package app

import (
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/stylebox"
)

type reactiveListAppImpl struct {
	components.Component
}

func New() ReactiveListApp {
	root := stylebox.New()
	return &reactiveListAppImpl{
		Component: nil,
	}
}
