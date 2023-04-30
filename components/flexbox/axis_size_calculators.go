package flexbox

import (
	"github.com/mieubrisse/box-layout-test/utilities"
)

type axisSizeCalculator func(
	desiredSizes []int,
	shouldGrow []bool,
	spaceAvailable int,
) axisSizeCalculationResults

type axisSizeCalculationResults struct {
	actualSizes []int

	spaceUsedByChildren int
}

// TODO move to be a function on the axis?
func calculateActualCrossAxisSizes(
	desiredSizes []int,
	shouldGrow []bool,
	// How much space is available in the cross axis
	spaceAvailable int,
) axisSizeCalculationResults {
	actualSizes := make([]int, len(desiredSizes))

	// The space used in the cross axis is the max across all children
	maxSpaceUsed := 0
	for idx, desiredSize := range desiredSizes {
		actualSize := desiredSize
		if shouldGrow[idx] {
			actualSize = utilities.GetMaxInt(actualSize, spaceAvailable)
		}

		// Ensure we don't overrun
		actualSize = utilities.GetMinInt(actualSize, spaceAvailable)

		actualSizes[idx] = actualSize
		maxSpaceUsed = utilities.GetMaxInt(actualSize, maxSpaceUsed)
	}
	return axisSizeCalculationResults{
		actualSizes:         actualSizes,
		spaceUsedByChildren: maxSpaceUsed,
	}
}

func calculateActualMainAxisSizes(
	desiredSizes []int,
	shouldGrow []bool,
	spaceAvailable int,
) axisSizeCalculationResults {
	totalDesiredSize := 0
	for _, desiredSize := range desiredSizes {
		totalDesiredSize += desiredSize
	}

	actualSizes := desiredSizes
	freeSpace := spaceAvailable - totalDesiredSize
	// The "grow" case
	if freeSpace > 0 {
		weights := make([]int, len(desiredSizes))
		for idx, desiredSize := range desiredSizes {
			if shouldGrow[idx] {
				// TODO deal with actual weights
				weights[idx] = desiredSize
				continue
			}

			weights[idx] = 0
		}

		actualSizes = utilities.DistributeSpaceByWeight(freeSpace, desiredSizes, weights)
		// The "shrink" case
	} else if freeSpace < 0 {
		// We use desired sizes as the weight, so that
		actualSizes = utilities.DistributeSpaceByWeight(freeSpace, desiredSizes, desiredSizes)
	}

	totalSpaceUsed := 0
	for _, spaceUsedByChild := range actualSizes {
		totalSpaceUsed += spaceUsedByChild
	}

	return axisSizeCalculationResults{
		actualSizes:         actualSizes,
		spaceUsedByChildren: totalSpaceUsed,
	}
}
