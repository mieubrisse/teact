package list

import (
	"github.com/mieubrisse/teact/teact/components"
	flexbox2 "github.com/mieubrisse/teact/teact/components/flexbox"
)

// Very simple container around a vertically-oriented flexbox
type List[T components.Component] interface {
	flexbox2.Flexbox

	GetItems() []T
	SetItems(items []T) List[T]
}
