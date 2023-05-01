package flexbox

import (
	"github.com/mieubrisse/teact/utilities"
)

type axisSizeCalculator func(
	minContentSizes []int,
	maxContentSizes []int,
	shouldGrow []bool,
	spaceAvailable int,
) axisSizeCalculationResults

type axisSizeCalculationResults struct {
	actualSizes []int

	spaceUsedByChildren int
}

// TODO move to be a function on the axis?
func calculateActualCrossAxisSizes(
	minContentSizes []int,
	maxContentSizes []int,
	shouldGrow []bool,
	// How much space is available in the cross axis
	spaceAvailable int,
) axisSizeCalculationResults {
	actualSizes := make([]int, len(maxContentSizes))

	// The space used in the cross axis is the max across all children
	maxSpaceUsed := 0
	for idx, max := range maxContentSizes {
		if shouldGrow[idx] {
			max = spaceAvailable
		}

		actualSize := utilities.GetMinInt(max, spaceAvailable)

		actualSizes[idx] = actualSize
		maxSpaceUsed = utilities.GetMaxInt(actualSize, maxSpaceUsed)
	}
	return axisSizeCalculationResults{
		actualSizes:         actualSizes,
		spaceUsedByChildren: maxSpaceUsed,
	}
}

func calculateActualMainAxisSizes(
	minContentSizes []int,
	maxContentSizes []int,
	shouldGrow []bool,
	spaceAvailable int,
) axisSizeCalculationResults {
	actualSizes := make([]int, len(minContentSizes))

	// First, allocate space from start_child to end_child, trying to getting each child to min-content before
	// proceeding to the next
	spaceUsedGettingToMin := 0
	for idx, min := range minContentSizes {
		sizeForItem := utilities.GetMinInt(min, spaceAvailable-spaceUsedGettingToMin)
		actualSizes[idx] = sizeForItem
		spaceUsedGettingToMin += sizeForItem
	}

	// If we used all the space attempting to get to min, we're done
	if spaceUsedGettingToMin == spaceAvailable {
		return axisSizeCalculationResults{
			actualSizes:         actualSizes,
			spaceUsedByChildren: spaceUsedGettingToMin,
		}
	}

	// We still have space, so start to allocate it amongst the items who can grow
	minDesiredSize := 0
	for _, min := range minContentSizes {
		minDesiredSize += min
	}
	maxDesiredSize := 0
	for _, max := range maxContentSizes {
		maxDesiredSize += max
	}
	spaceForGettingToMaxDesired := utilities.GetMinInt(spaceAvailable, maxDesiredSize) - minDesiredSize

	weightsForGettingToMax := make([]int, len(maxContentSizes))
	for idx, max := range maxContentSizes {
		min := minContentSizes[idx]

		// Each item gets a proportion of the space weighted by how far they are from their max
		weightsForGettingToMax[idx] = max - min
	}
	actualSizes = utilities.DistributeSpaceByWeight(spaceForGettingToMaxDesired, actualSizes, weightsForGettingToMax)

	// If we used all the space attempting to get to max, we're done
	spaceUsedGettingToMax := 0
	for _, size := range actualSizes {
		spaceUsedGettingToMax += size
	}
	if spaceUsedGettingToMax == spaceAvailable {
		return axisSizeCalculationResults{
			actualSizes:         actualSizes,
			spaceUsedByChildren: spaceUsedGettingToMax,
		}
	}

	// At this point, we *still* have space left over so give it to the children who can grow
	spaceForFillingBox := spaceAvailable - maxDesiredSize
	weightsForFillingBox := make([]int, len(maxContentSizes))
	for idx, itemShouldGrow := range shouldGrow {
		if itemShouldGrow {
			// TODO actual weights
			weightsForFillingBox[idx] = 1
		}
	}

	actualSizes = utilities.DistributeSpaceByWeight(spaceForFillingBox, actualSizes, weightsForFillingBox)

	totalSizeUsedByChildren := 0
	for _, size := range actualSizes {
		totalSizeUsedByChildren += size
	}

	return axisSizeCalculationResults{
		actualSizes:         actualSizes,
		spaceUsedByChildren: totalSizeUsedByChildren,
	}
}
