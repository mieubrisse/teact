package highlightable_list

import (
	"github.com/mieubrisse/teact/components/list"
	"github.com/mieubrisse/teact/utilities"
)

type HighlightableList[T HighlightableComponent] interface {
	list.List[T]

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

func (i impl[T]) Scroll(offset int) HighlightableList[T] {
	newHighlightedIdx := utilities.Clamp(i.highlightedIdx, 0, len(i.List.GetItems())-1)
	if i.highlightedIdx == newHighlightedIdx {
		return i
	}

	items := i.List.GetItems()
	if len(items) == 0 {
		return i
	}

	items[i.highlightedIdx].SetHighlight(false)
	items[newHighlightedIdx].SetHighlight(true)

	i.highlightedIdx = newHighlightedIdx
	return i
}
