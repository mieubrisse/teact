package highlightable_list

import (
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/list"
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
