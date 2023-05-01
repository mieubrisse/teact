package highlightable_list

import (
	"github.com/mieubrisse/teact/teact/components"
)

type HighlightableComponent interface {
	components.Component

	IsHighlighted() bool
	SetHighlight(isHighlighted bool) HighlightableComponent
}
