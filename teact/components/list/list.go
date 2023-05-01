package list

import (
	"github.com/mieubrisse/teact/teact/components"
	flexbox2 "github.com/mieubrisse/teact/teact/components/flexbox"
	"github.com/mieubrisse/teact/teact/components/flexbox_item"
)

// Very simple container around a vertically-oriented flexbox
type List[T components.Component] interface {
	flexbox2.Flexbox

	GetItems() []T
	SetItems(items []T) List[T]
}

type impl[T components.Component] struct {
	flexbox2.Flexbox

	items []T
}

func New[T components.Component]() List[T] {
	root := flexbox2.New().SetDirection(flexbox2.Column)
	return &impl[T]{
		items:   []T{},
		Flexbox: root,
	}
}

func NewWithContents[T components.Component](contents ...T) List[T] {
	elem := New[T]()
	elem.SetItems(contents)
	return elem
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