package highlightable_list

import (
	"github.com/mieubrisse/teact/components/list"
	"github.com/mieubrisse/teact/utilities"
)

type HighlightableList[T HighlightableComponent] interface {
	list.List[T]

	GetHighlightedIdx() int
	SetHighlightedIdx() int
	// Scrolls the highlighted item, with safeguards to prevent scrolling off the end of the list
	Scroll(offset int) HighlightableList[T]
}

type impl[T HighlightableComponent] struct {
	list.List[T]

	highlightedIdx int
}

func New[T HighlightableComponent]() HighlightableList[T] {
	return &impl[T]{
		List: list.New[T](),
	}
}

func (i *impl[T]) GetHighlightedIdx() int {
	return i.highlightedIdx
}

func (i *impl[T]) SetHighlightedIdx(newIdx int) HighlightableList[T] {
	if i.highlightedIdx == newIdx {
		return i
	}

	items := i.List.GetItems()
	if len(items) == 0 {
		return i
	}

	items[i.highlightedIdx].SetHighlight(false)
	items[newIdx].SetHighlight(true)
	i.highlightedIdx = newIdx
	return i
}

func (i *impl[T]) Scroll(offset int) HighlightableList[T] {
	newIdx := utilities.Clamp(i.highlightedIdx, 0, len(i.List.GetItems())-1)
	i.SetHighlightedIdx(newIdx)
	return i
}

func (i *impl[T]) SetItems(newItems []T) list.List[T] {
	items := i.List.GetItems()
	if len(items) > 0 {
		items[i.highlightedIdx].SetHighlight(false)
	}
	i.List.SetItems(newItems)

	i.highlightedIdx = 0
	if len(newItems) > 0 {
		newItems[i.highlightedIdx].SetHighlight(true)
	}

	return i
}
