package flexbox_item

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/components"
)

type OverflowStyle int

const (
	Wrap OverflowStyle = iota
	Truncate
)

// These are simply conveniences for the flexbox.NewWithContent , so that it's super easy to declare a single-item box
type FlexboxItemOpt func(item FlexboxItem)

func WithMinWidth(min FlexboxItemDimensionValue) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetMinWidth(min)
	}
}

func WithMaxWidth(max FlexboxItemDimensionValue) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetMaxWidth(max)
	}
}

func WithMinHeight(min FlexboxItemDimensionValue) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetMinHeight(min)
	}
}

func WithMaxHeight(max FlexboxItemDimensionValue) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetMaxHeight(max)
	}
}

func WithOverflowStyle(style OverflowStyle) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetOverflowStyle(style)
	}
}

func WithHorizontalGrowthFactor(growthFactor int) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetHorizontalGrowthFactor(growthFactor)
	}
}

func WithVerticalGrowthFactor(growthFactor int) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetVerticalGrowthFactor(growthFactor)
	}
}

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

type flexboxItemImpl struct {
	component components.Component

	// These determine how the item flexes
	// This is analogous to both "flex-basis" and "flex-grow", where:
	// - MaxAvailable indicates "flex-grow: >1" (see weight below)
	// - Anything else indicates "flex-grow: 0", and sets the "flex-basis"
	minWidth  FlexboxItemDimensionValue
	maxWidth  FlexboxItemDimensionValue
	minHeight FlexboxItemDimensionValue
	maxHeight FlexboxItemDimensionValue

	overflowStyle OverflowStyle

	// Analogous to "flex-grow" when on the main axis, and "align-items: stretch" when on the cross axis (on a per-item basis)
	horizontalGrowthFactor int

	// Analogous to "flex-grow" when on the main axis, and "align-items: stretch" when on the cross axis (on a per-item basis)
	verticalGrowthFactor int
}

// TODO add varargs Opts to make it easier to adjust
func New(component components.Component, opts ...FlexboxItemOpt) FlexboxItem {
	result := &flexboxItemImpl{
		component: component,
		// TODO move a lot of these out into its own class????
		minWidth:               MinContent,
		maxWidth:               MaxContent,
		minHeight:              MinContent,
		maxHeight:              MaxContent,
		overflowStyle:          Wrap,
		horizontalGrowthFactor: 0,
		verticalGrowthFactor:   0,
	}

	for _, opt := range opts {
		opt(result)
	}
	return result
}

func (item *flexboxItemImpl) GetContentMinMax() (minWidth int, maxWidth int, minHeight int, maxHeight int) {
	innerMinWidth, innerMaxWidth, innerMinHeight, innerMaxHeight := item.GetComponent().GetContentMinMax()
	itemMinWidth, itemMaxWidth, itemMinHeight, itemMaxHeight := calculateFlexboxItemContentSizesFromInnerContentSizes(
		innerMinWidth,
		innerMaxWidth,
		innerMinHeight,
		innerMaxHeight,
		item,
	)

	return itemMinWidth, itemMaxWidth, itemMinHeight, itemMaxHeight
}

func (item *flexboxItemImpl) SetWidthAndGetDesiredHeight(width int) int {
	// TODO we're redoing this calculation when we've already done it - if we cache it, we'll save extra work
	return item.component.SetWidthAndGetDesiredHeight(width)
}

func (item *flexboxItemImpl) View(width int, height int) string {
	if width == 0 || height == 0 {
		return ""
	}

	component := item.GetComponent()

	var widthWhenRendering int
	switch item.GetOverflowStyle() {
	case Wrap:
		widthWhenRendering = width
	case Truncate:
		// If truncating, the child will _think_ they have infinite space available
		// and then we'll truncate them later
		_, maxWidth, _, _ := component.GetContentMinMax()
		widthWhenRendering = maxWidth
	default:
		panic(fmt.Sprintf("Unknown item overflow style: %v", item.GetOverflowStyle()))
	}

	// TODO allow column format
	result := component.View(widthWhenRendering, height)

	// Truncate, in case the inner item runs over (which will almost definitely be the case when overflowStyle = Truncate)
	result = lipgloss.NewStyle().
		MaxWidth(width).
		MaxHeight(height).
		Render(result)

	// Now expand, in case the inner item is smaller than what we need
	result = lipgloss.NewStyle().
		Width(width).
		Height(height).
		Render(result)

	return result
}

func (item *flexboxItemImpl) GetComponent() components.Component {
	return item.component
}

func (item *flexboxItemImpl) GetMinWidth() FlexboxItemDimensionValue {
	return item.minWidth
}

func (item *flexboxItemImpl) SetMinWidth(min FlexboxItemDimensionValue) FlexboxItem {
	item.minWidth = min
	return item
}

func (item *flexboxItemImpl) GetMaxWidth() FlexboxItemDimensionValue {
	return item.maxWidth
}

func (item *flexboxItemImpl) SetMaxWidth(max FlexboxItemDimensionValue) FlexboxItem {
	item.maxWidth = max
	return item
}

func (item *flexboxItemImpl) GetMinHeight() FlexboxItemDimensionValue {
	return item.minHeight
}

func (item *flexboxItemImpl) SetMinHeight(min FlexboxItemDimensionValue) FlexboxItem {
	item.minHeight = min
	return item
}

func (item *flexboxItemImpl) GetMaxHeight() FlexboxItemDimensionValue {
	return item.maxHeight
}

func (item *flexboxItemImpl) SetMaxHeight(max FlexboxItemDimensionValue) FlexboxItem {
	item.maxHeight = max
	return item
}

func (item *flexboxItemImpl) GetOverflowStyle() OverflowStyle {
	return item.overflowStyle
}

func (item *flexboxItemImpl) SetOverflowStyle(style OverflowStyle) FlexboxItem {
	item.overflowStyle = style
	return item
}

func (item *flexboxItemImpl) GetHorizontalGrowthFactor() int {
	return item.horizontalGrowthFactor
}

func (item *flexboxItemImpl) SetHorizontalGrowthFactor(growFactor int) FlexboxItem {
	item.horizontalGrowthFactor = growFactor
	return item
}

func (item *flexboxItemImpl) GetVerticalGrowthFactor() int {
	return item.verticalGrowthFactor
}

func (item *flexboxItemImpl) SetVerticalGrowthFactor(growthFactor int) FlexboxItem {
	item.verticalGrowthFactor = growthFactor
	return item
}

// ====================================================================================================
//                                   Private Helper Functions
// ====================================================================================================

// Rescales an item's content size based on the per-item configuration the user has set
// Max is guaranteed to be >= min
func calculateFlexboxItemContentSizesFromInnerContentSizes(
	innerMinWidth,
	innertMaxWidth,
	innerMinHeight,
	innerMaxHeight int,
	item FlexboxItem,
) (itemMinWidth, itemMaxWidth, itemMinHeight, itemMaxHeight int) {
	itemMinWidth = item.GetMinWidth().getSizeRetriever()(innerMinWidth, innertMaxWidth)
	itemMaxWidth = item.GetMaxWidth().getSizeRetriever()(innerMinWidth, innertMaxWidth)

	if itemMaxWidth < itemMinWidth {
		itemMaxWidth = itemMinWidth
	}

	// TODO there's a very minor bug here where if we use a fixed-size width, the height min-content should go down
	//  but it doesn't because we don't recalculate the actual height based on the actual width
	//  The way to fix this is to figure out how extrinsic width/height settings (e.g. 60px, 20%, etc.) can be factored
	//  into our calculations

	itemMinHeight = item.GetMinHeight().getSizeRetriever()(innerMinHeight, innerMaxHeight)
	itemMaxHeight = item.GetMaxHeight().getSizeRetriever()(innerMinHeight, innerMaxHeight)

	if itemMaxHeight < itemMinHeight {
		itemMaxHeight = itemMinHeight
	}

	return
}
