package list

import (
	"github.com/mieubrisse/teact/components"
	"github.com/mieubrisse/teact/components/flexbox"
	"github.com/mieubrisse/teact/components/flexbox_item"
)

// Very simple container around a vertically-oriented flexbox
type List[T components.Component] interface {
	components.Component

	GetItems() []T
	SetItems(items []T) List[T]
}

type impl[T components.Component] struct {
	items []T

	root flexbox.Flexbox
}

func New[T components.Component]() List[T] {
	root := flexbox.New().SetDirection(flexbox.Column)
	return &impl[T]{
		items: []T{},
		root:  root,
	}
}

func (i impl[T]) GetItems() []T {
	return i.items
}

func (i *impl[T]) SetItems(items []T) List[T] {
	i.items = items

	flexboxItems := make([]flexbox_item.FlexboxItem, len(items))
	for idx, item := range items {
		flexboxItems[idx] = flexbox_item.New(item).SetMaxWidth(flexbox_item.MaxAvailable)
	}
	i.root.SetChildren(flexboxItems)

	return i
}

func (i impl[T]) GetContentMinMax() (minWidth, maxWidth, minHeight, maxHeight int) {
	return i.root.GetContentMinMax()
}

func (i impl[T]) GetContentHeightForGivenWidth(width int) int {
	return i.root.GetContentHeightForGivenWidth(width)
}

func (i impl[T]) View(width int, height int) string {
	return i.root.View(width, height)
}
