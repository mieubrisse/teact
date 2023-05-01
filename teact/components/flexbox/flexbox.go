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

type minMaxChildDimensionsCache struct {
	minWidths  []int
	maxWidths  []int
	minHeights []int
	maxHeights []int
}

type flexboxImpl struct {
	children []flexbox_item.FlexboxItem

	direction Direction

	horizontalAlignment AxisAlignment
	verticalAlignment   AxisAlignment

	// -------------------- Calculation Caching -----------------------
	// The min/max widths/heights of children
	childDimensionsCache minMaxChildDimensionsCache

	// The actual widths each child will get (cached between GetContentHeightForGivenWidth and View)
	actualChildWidthsCache axisSizeCalculationResults

	// The desired height each child wants given its width (cached between GetContentHeightForGivenWidth and View)
	desiredChildHeightsGivenWidthCache []int
}

func New(items ...flexbox_item.FlexboxItem) Flexbox {
	return &flexboxImpl{
		children:                           items,
		direction:                          Row,
		horizontalAlignment:                AlignStart,
		verticalAlignment:                  AlignStart,
		actualChildWidthsCache:             axisSizeCalculationResults{},
		desiredChildHeightsGivenWidthCache: nil,
	}
}

func NewWithOpts(items []flexbox_item.FlexboxItem, opts ...FlexboxOpt) Flexbox {
	result := New(items...)
	for _, opt := range opts {
		opt(result)
	}
	return result
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

	// Cache, so that future steps don't need to recalculate this
	b.childDimensionsCache = minMaxChildDimensionsCache{
		minWidths:  childMinWidths,
		maxWidths:  childMaxWidths,
		minHeights: childMinHeights,
		maxHeights: childMaxHeights,
	}

	minWidth := b.direction.reduceChildWidths(childMinWidths)
	maxWidth := b.direction.reduceChildWidths(childMaxWidths)

	minHeight := b.direction.reduceChildHeights(childMinHeights)
	maxHeight := b.direction.reduceChildHeights(childMaxHeights)

	return minWidth, maxWidth, minHeight, maxHeight
}

func (b *flexboxImpl) SetWidthAndGetDesiredHeight(width int) int {
	if width == 0 {
		return 0
	}

	// Width
	growthFactors := make([]int, len(b.children))
	for idx, item := range b.children {
		growthFactors[idx] = item.GetHorizontalGrowthFactor()
	}
	actualWidthsCalcResults := b.direction.getActualWidths(
		b.childDimensionsCache.minWidths,
		b.childDimensionsCache.maxWidths,
		growthFactors,
		width,
	)

	// Cache the result, so we don't have to recalculate it in View
	b.actualChildWidthsCache = actualWidthsCalcResults

	desiredHeights := make([]int, len(b.children))
	for idx, item := range b.children {
		actualWidth := actualWidthsCalcResults.actualSizes[idx]
		desiredHeight := item.SetWidthAndGetDesiredHeight(actualWidth)

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

	// TODO get rid of this.. doesn't quite make sense
	growthFactors := make([]int, len(b.children))
	for idx, item := range b.children {
		growthFactors[idx] = item.GetVerticalGrowthFactor()
	}
	actualHeightsCalcResult := b.direction.getActualHeights(
		b.childDimensionsCache.minHeights,
		b.desiredChildHeightsGivenWidthCache,
		growthFactors,
		height,
	)

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
