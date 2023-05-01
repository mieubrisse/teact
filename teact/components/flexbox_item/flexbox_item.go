package flexbox_item

import (
	"github.com/mieubrisse/teact/teact/components"
)

type OverflowStyle int

const (
	Wrap OverflowStyle = iota
	Truncate
)

type FlexboxItem interface {
	components.Component

	GetComponent() components.Component

	GetMinWidth() FlexboxItemDimensionValue
	SetMinWidth(min FlexboxItemDimensionValue) FlexboxItem
	GetMaxWidth() FlexboxItemDimensionValue
	SetMaxWidth(max FlexboxItemDimensionValue) FlexboxItem

	GetMinHeight() FlexboxItemDimensionValue
	SetMinHeight(min FlexboxItemDimensionValue) FlexboxItem
	GetMaxHeight() FlexboxItemDimensionValue
	SetMaxHeight(max FlexboxItemDimensionValue) FlexboxItem

	GetOverflowStyle() OverflowStyle
	SetOverflowStyle(style OverflowStyle) FlexboxItem

	// Analogous to "flex-grow" when on the main axis, and "align-items: stretch" when on the cross axis (on a per-item basis)
	// 0 means no growth
	GetHorizontalGrowthFactor() int
	SetHorizontalGrowthFactor(growthFactor int) FlexboxItem
	GetVerticalGrowthFactor() int
	SetVerticalGrowthFactor(growthFactor int) FlexboxItem
}
