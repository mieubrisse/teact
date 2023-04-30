package flexbox

import (
	"github.com/charmbracelet/lipgloss"
)

// The direction that the flexbox ought to be layed out in
type Direction interface {
	reduceChildWidths(childWidths []int) int

	reduceChildHeights(childHeights []int) int

	getActualWidths(desiredWidths []int, shouldGrow []bool, widthAvailable int) axisSizeCalculationResults

	getActualHeights(desiredHeights []int, shouldGrow []bool, heightAvailable int) axisSizeCalculationResults

	renderContentFragments(contentFragments []string, width int, height int, horizontalAlignment AxisAlignment, verticalAlignment AxisAlignment) string
}

// Row lays out the flexbox items in a row, left to right
// The flex direction will be horizontal
// Corresponds to "flex-direction: row" in CSS
var Row = &directionImpl{
	actualWidthCalculator:  calculateActualMainAxisSizes,
	actualHeightCalculator: calculateActualCrossAxisSizes,
	widthDimensionReducer:  mainAxisDimensionReducer,
	heightDimensionReducer: crossAxisDimensionReducer,
	contentFragmentRenderer: func(contentFragments []string, width int, height int, horizontalAlign AxisAlignment, verticalAlign AxisAlignment) string {
		joined := lipgloss.JoinHorizontal(lipgloss.Position(verticalAlign), contentFragments...)
		horizontallyPlaced := lipgloss.PlaceHorizontal(width, lipgloss.Position(horizontalAlign), joined)
		return lipgloss.PlaceVertical(height, lipgloss.Position(verticalAlign), horizontallyPlaced)
	},
}

// Column lays out the flexbox items in a column, top to bottom
// The flex direction will be vertical
// Corresponds to "flex-direction: column" in CSS
var Column = &directionImpl{
	actualWidthCalculator:  calculateActualCrossAxisSizes,
	actualHeightCalculator: calculateActualMainAxisSizes,
	widthDimensionReducer:  crossAxisDimensionReducer,
	heightDimensionReducer: mainAxisDimensionReducer,
	contentFragmentRenderer: func(contentFragments []string, width int, height int, horizontalAlign AxisAlignment, verticalAlign AxisAlignment) string {
		joined := lipgloss.JoinVertical(lipgloss.Position(horizontalAlign), contentFragments...)
		horizontallyPlaced := lipgloss.PlaceHorizontal(width, lipgloss.Position(horizontalAlign), joined)
		return lipgloss.PlaceVertical(height, lipgloss.Position(verticalAlign), horizontallyPlaced)
	},
}

// ====================================================================================================
//
//	Private
//
// ====================================================================================================
type directionImpl struct {
	actualWidthCalculator   axisSizeCalculator
	actualHeightCalculator  axisSizeCalculator
	widthDimensionReducer   axisDimensionMinMaxCombiner
	heightDimensionReducer  axisDimensionMinMaxCombiner
	contentFragmentRenderer func(contentFragments []string, width int, height int, horizontalAlign AxisAlignment, verticalAlign AxisAlignment) string
}

func (a directionImpl) reduceChildWidths(childWidths []int) int {
	return a.widthDimensionReducer(childWidths)
}

func (a directionImpl) reduceChildHeights(childHeights []int) int {
	return a.heightDimensionReducer(childHeights)
}

func (r directionImpl) getActualWidths(desiredWidths []int, shouldGrow []bool, widthAvailable int) axisSizeCalculationResults {
	return r.actualWidthCalculator(
		desiredWidths,
		shouldGrow,
		widthAvailable,
	)
}

func (r directionImpl) getActualHeights(desiredHeights []int, shouldGrow []bool, heightAvailable int) axisSizeCalculationResults {
	return r.actualHeightCalculator(
		desiredHeights,
		shouldGrow,
		heightAvailable,
	)
}

func (r directionImpl) renderContentFragments(contentFragments []string, width int, height int, horizontalAlign AxisAlignment, verticalAlign AxisAlignment) string {
	return r.contentFragmentRenderer(contentFragments, width, height, horizontalAlign, verticalAlign)
}
