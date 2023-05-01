package flexbox

import (
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox_item"
)

// NOTE: This class does some stateful caching, so when you're testing methods like "View" make sure you call the
// full flow of GetContentMinMax -> SetWidthAndGetDesiredHeight -> View as necessary

type Flexbox interface {
	components.Component

	GetChildren() []flexbox_item.FlexboxItem
	SetChildren(children []flexbox_item.FlexboxItem) Flexbox

	GetDirection() Direction
	SetDirection(direction Direction) Flexbox

	GetHorizontalAlignment() AxisAlignment
	SetHorizontalAlignment(alignment AxisAlignment) Flexbox

	GetVerticalAlignment() AxisAlignment
	SetVerticalAlignment(alignment AxisAlignment) Flexbox
}
