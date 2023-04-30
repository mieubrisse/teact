package flexbox

import "github.com/mieubrisse/box-layout-test/utilities"

// Reduces all the child values for a given dimension to a single value for the parent
// Used for calculating what the flexbox's min width is based off child min widths, flexbox max height based off childrens', etc.
type axisDimensionMinMaxCombiner func(values []int) int

// Reduces all the childrens' cross axis dimension values into one for the flexbox parent
// The cross axis uses the max (because elements don't flex on the cross axis, so whichever is biggest dominates)
func crossAxisDimensionReducer(crossAxisDimensionValues []int) int {
	max := 0
	for _, value := range crossAxisDimensionValues {
		max = utilities.GetMaxInt(max, value)
	}
	return max
}

// Reduces all the childrens' main axis dimension values into one for the flexbox parent
// The main axis uses the sum, because elements flex on the main axis so the parent size is the size of all the childrens'
// main axis values
func mainAxisDimensionReducer(mainAxisDimensionValues []int) int {
	sum := 0
	for _, value := range mainAxisDimensionValues {
		sum += value
	}
	return sum
}
