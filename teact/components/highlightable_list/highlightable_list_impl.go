package highlightable_list

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/teact/components/list"
	"github.com/mieubrisse/teact/teact/utilities"
)

type highlightableListImpl[T HighlightableComponent] struct {
	list.List[T]

	highlightedIdx int

	isFocused bool
}

func New[T HighlightableComponent]() HighlightableList[T] {
	return &highlightableListImpl[T]{
		List: list.New[T](),
	}
}

func (impl *highlightableListImpl[T]) GetHighlightedIdx() int {
	return impl.highlightedIdx
}

func (impl *highlightableListImpl[T]) SetHighlightedIdx(newIdx int) HighlightableList[T] {
	if impl.highlightedIdx == newIdx {
		return impl
	}

	items := impl.List.GetItems()
	if len(items) == 0 {
		return impl
	}

	items[impl.highlightedIdx].SetHighlight(false)
	items[newIdx].SetHighlight(true)
	impl.highlightedIdx = newIdx
	return impl
}

func (impl *highlightableListImpl[T]) Scroll(offset int) HighlightableList[T] {
	newIdx := utilities.Clamp(impl.highlightedIdx+offset, 0, len(impl.List.GetItems())-1)
	impl.SetHighlightedIdx(newIdx)
	return impl
}

func (impl *highlightableListImpl[T]) SetItems(newItems []T) list.List[T] {
	items := impl.List.GetItems()
	if len(items) > 0 {
		items[impl.highlightedIdx].SetHighlight(false)
	}
	impl.List.SetItems(newItems)

	impl.highlightedIdx = 0
	if len(newItems) > 0 {
		newItems[impl.highlightedIdx].SetHighlight(true)
	}

	return impl
}

func (impl *highlightableListImpl[T]) Update(msg tea.Msg) tea.Cmd {
	if !impl.isFocused {
		return nil
	}

	// TDOO extract this to a helper (or embedded struct?) or something
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+j", "down":
			impl.Scroll(1)
			return nil
		case "ctrl+k", "up":
			impl.Scroll(-1)
			return nil
		}
	}

	items := impl.GetItems()
	if len(items) == 0 {
		return nil
	}

	return utilities.TryUpdate(items[impl.highlightedIdx], msg)
}

func (impl *highlightableListImpl[T]) SetFocus(isFocused bool) tea.Cmd {
	impl.isFocused = isFocused

	items := impl.GetItems()
	if len(items) == 0 {
		return nil
	}

	items[impl.highlightedIdx].SetHighlight(isFocused)
	return nil
}

func (impl *highlightableListImpl[T]) IsFocused() bool {
	return impl.isFocused
}
