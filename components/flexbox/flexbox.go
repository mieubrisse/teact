package flexbox

import (
	"github.com/mieubrisse/box-layout-test/components"
	"github.com/mieubrisse/box-layout-test/components/flexbox_item"
)

// NOTE: This class does some stateful caching, so when you're testing methods like "View" make sure you call the
// full flow of GetContentMinMax -> GetContentHeightForGivenWidth -> View as necessary

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

type flexboxImpl struct {
	children []flexbox_item.FlexboxItem

	direction Direction

	horizontalAlignment AxisAlignment
	verticalAlignment   AxisAlignment

	// -------------------- Calculation Caching -----------------------
	// The actual widths each child will get (cached between GetContentHeightForGivenWidth and View)
	actualChildWidthsCache axisSizeCalculationResults

	// The desired height each child wants given its width (cached between GetContentHeightForGivenWidth and View)
	desiredChildHeightsGivenWidthCache []int
}

// Convenience constructor for a box with a single element
func NewWithContent(component components.Component, opts ...flexbox_item.FlexboxItemOpt) Flexbox {
	item := flexbox_item.New(component)
	for _, opt := range opts {
		opt(item)
	}
	return NewWithContents(item)
}

// Convenience constructor for a box with multiple elements
func NewWithContents(items ...flexbox_item.FlexboxItem) Flexbox {
	return New().SetChildren(items)
}

func New() Flexbox {
	return &flexboxImpl{
		children:                           make([]flexbox_item.FlexboxItem, 0),
		direction:                          Row,
		horizontalAlignment:                AlignStart,
		verticalAlignment:                  AlignStart,
		actualChildWidthsCache:             axisSizeCalculationResults{},
		desiredChildHeightsGivenWidthCache: nil,
	}
}

func (b *flexboxImpl) GetChildren() []flexbox_item.FlexboxItem {
	return b.children
}

func (b *flexboxImpl) SetChildren(children []flexbox_item.FlexboxItem) Flexbox {
	b.children = children
	return b
}

func (b *flexboxImpl) GetDirection() Direction {
	return b.direction
}

func (b *flexboxImpl) SetDirection(direction Direction) Flexbox {
	b.direction = direction
	return b
}

func (b flexboxImpl) GetHorizontalAlignment() AxisAlignment {
	return b.horizontalAlignment
}

func (b *flexboxImpl) SetHorizontalAlignment(alignment AxisAlignment) Flexbox {
	b.horizontalAlignment = alignment
	return b
}

func (b flexboxImpl) GetVerticalAlignment() AxisAlignment {
	return b.verticalAlignment
}

func (b *flexboxImpl) SetVerticalAlignment(alignment AxisAlignment) Flexbox {
	b.verticalAlignment = alignment
	return b
}

func (b *flexboxImpl) GetContentMinMax() (int, int, int, int) {
	numChildren := len(b.children)
	childMinWidths := make([]int, numChildren)
	childMaxWidths := make([]int, numChildren)
	childMinHeights := make([]int, numChildren)
	childMaxHeights := make([]int, numChildren)
	for idx, item := range b.children {
		childMinWidths[idx], childMaxWidths[idx], childMinHeights[idx], childMaxHeights[idx] = item.GetContentMinMax()
	}

	minWidth := b.direction.reduceChildWidths(childMinWidths)
	maxWidth := b.direction.reduceChildWidths(childMaxWidths)

	minHeight := b.direction.reduceChildHeights(childMinHeights)
	maxHeight := b.direction.reduceChildHeights(childMaxHeights)

	return minWidth, maxWidth, minHeight, maxHeight
}

func (b *flexboxImpl) GetContentHeightForGivenWidth(width int) int {
	if width == 0 {
		return 0
	}

	// Width
	desiredChildWidths := make([]int, len(b.children)) // NOTE: we actually already calculated this above, with GetContentMinMax. Maybe cache?
	shouldGrowWidths := make([]bool, len(b.children))
	for idx, item := range b.children {
		_, desiredChildWidths[idx], _, _ = item.GetComponent().GetContentMinMax()
		shouldGrowWidths[idx] = item.GetMaxWidth().ShouldGrow()
	}
	actualWidthsCalcResults := b.direction.getActualWidths(desiredChildWidths, shouldGrowWidths, width)

	// Cache the result, so we don't have to recalculate it in View
	b.actualChildWidthsCache = actualWidthsCalcResults

	desiredHeights := make([]int, len(b.children))
	for idx, item := range b.children {
		actualWidth := actualWidthsCalcResults.actualSizes[idx]
		desiredHeight := item.GetContentHeightForGivenWidth(actualWidth)

		desiredHeights[idx] = desiredHeight
	}

	// Cache the result, so we don't have to recalculate it in View
	b.desiredChildHeightsGivenWidthCache = desiredHeights

	return b.direction.reduceChildHeights(desiredHeights)
}

func (b *flexboxImpl) View(width int, height int) string {
	if width == 0 || height == 0 {
		return ""
	}

	actualWidths := b.actualChildWidthsCache.actualSizes
	// widthNotUsedByChildren := utilities.GetMaxInt(0, width-b.actualChildWidthsCache.spaceUsedByChildren)

	shouldGrowHeights := make([]bool, len(b.children))
	for idx, item := range b.children {
		shouldGrowHeights[idx] = item.GetMaxHeight().ShouldGrow()
	}
	actualHeightsCalcResult := b.direction.getActualHeights(b.desiredChildHeightsGivenWidthCache, shouldGrowHeights, height)

	actualHeights := actualHeightsCalcResult.actualSizes
	// heightNotUsedByChildren := utilities.GetMaxInt(0, height-actualHeightsCalcResult.spaceUsedByChildren)

	// Now render each child
	allContentFragments := make([]string, len(b.children))
	for idx, item := range b.children {
		childWidth := actualWidths[idx]
		childHeight := actualHeights[idx]
		childStr := item.View(childWidth, childHeight)

		allContentFragments[idx] = childStr
	}

	content := b.direction.renderContentFragments(allContentFragments, width, height, b.horizontalAlignment, b.verticalAlignment)

	return content
}
