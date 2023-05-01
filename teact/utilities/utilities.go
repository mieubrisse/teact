package utilities

func GetMaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetMinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Clamp(value, low, high int) int {
	if high < low {
		low, high = high, low
	}
	return GetMinInt(high, GetMaxInt(low, value))
}

// Distributes the space (which can be negative) across the children, using the weight as a bias for how to allocate
// The only scenario where no space will be distributed is if there is no total weight
// If the space does get distributed, it's guaranteed to be done exactly (no more or less will remain)
func DistributeSpaceByWeight(spaceToAllocate int, inputSizes []int, weights []int) []int {
	result := make([]int, len(inputSizes))
	for idx, inputSize := range inputSizes {
		result[idx] = inputSize
	}

	totalWeight := 0
	for _, weight := range weights {
		totalWeight += weight
	}

	// watch out for divide-by-zero
	if totalWeight == 0 {
		return result
	}

	desiredSpaceAllocated := float64(0)
	actualSpaceAllocated := 0
	for idx, size := range inputSizes {
		result[idx] = size

		// Dump any remaining space for the last item (it should always be at most 1
		// in any direction)
		if idx == len(inputSizes)-1 {
			result[idx] += spaceToAllocate - actualSpaceAllocated
			break
		}

		weight := weights[idx]
		share := float64(weight) / float64(totalWeight)

		// Because we can only display lines in integer numbers, but flexing
		// will yield float scale ratios, no matter what space we give each item
		// our integer value will always be off from the float value
		// This algorithm is to ensure that we're always rounding in the direction
		// that pushes us closer to our desired allocation (rather than naively rounding up or down)
		desiredSpaceForItem := share * float64(spaceToAllocate)
		var actualSpaceForItem int
		if desiredSpaceAllocated < float64(actualSpaceAllocated) {
			// If we're under our desired allocation, round up to try and get closer
			actualSpaceForItem = int(desiredSpaceForItem + 1)
		} else {
			// If we're at or over our desired allocation, round down (so we either stay there or get closer by undershooting)
			actualSpaceForItem = int(desiredSpaceForItem)
		}

		result[idx] += actualSpaceForItem
		desiredSpaceAllocated += desiredSpaceForItem
		actualSpaceAllocated += actualSpaceForItem
	}

	return result
}
