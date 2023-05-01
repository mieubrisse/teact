package list

import (
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox"
)

// Very simple container around a vertically-oriented flexbox
type List[T components.Component] interface {
	flexbox.Flexbox

	GetItems() []T
	SetItems(items []T) List[T]
}
