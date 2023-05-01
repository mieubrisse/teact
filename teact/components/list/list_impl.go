package list

import (
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox"
	"github.com/mieubrisse/teact/teact/components/flexbox_item"
)

type impl[T components.Component] struct {
	flexbox.Flexbox

	items []T
}

func New[T components.Component]() List[T] {
	root := flexbox.New().SetDirection(flexbox.Column)
	return &impl[T]{
		items:   []T{},
		Flexbox: root,
	}
}

func NewWithContents[T components.Component](contents ...T) List[T] {
	root := flexbox.New().SetDirection(flexbox.Column)
	return &impl[T]{
		items:   contents,
		Flexbox: root,
	}
}

func (i impl[T]) GetItems() []T {
	return i.items
}

func (i *impl[T]) SetItems(items []T) List[T] {
	i.items = items

	flexboxItems := make([]flexbox_item.FlexboxItem, len(items))
	for idx, item := range items {
		flexboxItems[idx] = flexbox_item.New(item).SetHorizontalGrowthFactor(1)
	}
	i.Flexbox.SetChildren(flexboxItems)

	return i
}

func (i impl[T]) GetContentMinMax() (minWidth, maxWidth, minHeight, maxHeight int) {
	return i.Flexbox.GetContentMinMax()
}

func (i impl[T]) SetWidthAndGetDesiredHeight(width int) int {
	return i.Flexbox.SetWidthAndGetDesiredHeight(width)
}

func (i impl[T]) View(width int, height int) string {
	return i.Flexbox.View(width, height)
}
