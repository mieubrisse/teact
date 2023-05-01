package highlightable_list

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/list"
	"github.com/mieubrisse/teact/teact/utilities"
)

type HighlightableList[T HighlightableComponent] interface {
	list.List[T]
	components.InteractiveComponent

	GetHighlightedIdx() int
	SetHighlightedIdx(idx int) HighlightableList[T]
	// Scrolls the highlighted item, with safeguards to prevent scrolling off the end of the list
	Scroll(offset int) HighlightableList[T]

	// TODO something about keeping items highlighted when losing focus
}

type impl[T HighlightableComponent] struct {
	list.List[T]

	highlightedIdx int

	isFocused bool
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
	newIdx := utilities.Clamp(i.highlightedIdx+offset, 0, len(i.List.GetItems())-1)
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

func (i *impl[T]) Update(msg tea.Msg) tea.Cmd {
	if !i.isFocused {
		return nil
	}

	// TDOO extract this to a helper (or embedded struct?) or something
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			i.Scroll(1)
		case "k", "up":
			i.Scroll(-1)
		}
	}
	return nil
}

func (i *impl[T]) SetFocus(isFocused bool) tea.Cmd {
	i.isFocused = isFocused

	items := i.GetItems()
	if len(items) == 0 {
		return nil
	}

	items[i.highlightedIdx].SetHighlight(isFocused)
	return nil
}

func (i *impl[T]) IsFocused() bool {
	return i.isFocused
}