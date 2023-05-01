package flexbox

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// ====================================================================================================
//
//	Cross Axis Tests
//
// ====================================================================================================
func TestCrossAxisTruncationNoGrowth(t *testing.T) {
	// These don't do anything for the cross axis
	minSizes := []int{
		3,
		3,
		3,
	}
	maxSizes := []int{
		10,
		5,
		7,
	}
	shouldGrow := []bool{
		false,
		false,
		false,
	}
	calcResult := calculateActualCrossAxisSizes(minSizes, maxSizes, shouldGrow, 6)

	require.Equal(t, 6, calcResult.spaceUsedByChildren)
	require.Equal(t, []int{6, 5, 6}, calcResult.actualSizes)
}

func TestCrossAxisTruncationWithGrowth(t *testing.T) {
	// These don't do anything for the cross axis
	minSizes := []int{
		3,
		3,
		3,
	}
	maxSizes := []int{
		10,
		5,
		7,
	}
	shouldGrow := []bool{
		true,
		true,
		true,
	}
	calcResult := calculateActualCrossAxisSizes(minSizes, maxSizes, shouldGrow, 6)

	require.Equal(t, 6, calcResult.spaceUsedByChildren)
	require.Equal(t, []int{6, 6, 6}, calcResult.actualSizes)
}

func TestCrossAxisExtraSpaceNoGrowth(t *testing.T) {
	// These don't do anything for the cross axis
	minSizes := []int{
		3,
		3,
		3,
	}
	maxSizes := []int{
		10,
		5,
		7,
	}
	shouldGrow := []bool{
		false,
		false,
		false,
	}
	calcResult := calculateActualCrossAxisSizes(minSizes, maxSizes, shouldGrow, 12)

	require.Equal(t, 10, calcResult.spaceUsedByChildren)
	require.Equal(t, []int{10, 5, 7}, calcResult.actualSizes)
}

func TestCrossAxisExtraSpaceWithGrowth(t *testing.T) {
	// These don't do anything for the cross axis
	minSizes := []int{
		3,
		3,
		3,
	}
	maxSizes := []int{
		10,
		5,
		7,
	}
	shouldGrow := []bool{
		true,
		true,
		true,
	}
	calcResult := calculateActualCrossAxisSizes(minSizes, maxSizes, shouldGrow, 12)

	require.Equal(t, 12, calcResult.spaceUsedByChildren)
	require.Equal(t, []int{12, 12, 12}, calcResult.actualSizes)
}

// ====================================================================================================
//
//	Main Axis Tests
//
// ====================================================================================================
func TestMainAxisTruncation(t *testing.T) {
	minSizes := []int{
		5,
		5,
		5,
	}
	maxSizes := []int{
		7,
		7,
		7,
	}
	shouldGrow := []bool{
		false,
		false,
		false,
	}
	calcResult := calculateActualMainAxisSizes(minSizes, maxSizes, shouldGrow, 7)

	require.Equal(t, 7, calcResult.spaceUsedByChildren)
	require.Equal(t, []int{5, 2, 0}, calcResult.actualSizes)
}

func TestMainAxisAllFixed(t *testing.T) {
	minSizes := []int{
		5,
		5,
		5,
	}
	maxSizes := []int{
		5,
		5,
		5,
	}
	shouldGrow := []bool{
		false,
		false,
		false,
	}
	calcResult := calculateActualMainAxisSizes(minSizes, maxSizes, shouldGrow, 20)

	require.Equal(t, 15, calcResult.spaceUsedByChildren)
	require.Equal(t, []int{5, 5, 5}, calcResult.actualSizes)
}

func TestMainAxisSomeFixed(t *testing.T) {
	minSizes := []int{
		5,
		5,
		5,
	}
	maxSizes := []int{
		5,
		10,
		5,
	}
	shouldGrow := []bool{
		false,
		false,
		false,
	}
	calcResult := calculateActualMainAxisSizes(minSizes, maxSizes, shouldGrow, 22)

	require.Equal(t, 20, calcResult.spaceUsedByChildren)
	require.Equal(t, []int{5, 10, 5}, calcResult.actualSizes)
}

func TestMainAxisEvenGrowth(t *testing.T) {
	minSizes := []int{
		5,
		5,
		5,
	}
	maxSizes := []int{
		8,
		8,
		8,
	}
	shouldGrow := []bool{
		false,
		false,
		false,
	}
	calcResult := calculateActualMainAxisSizes(minSizes, maxSizes, shouldGrow, 18)

	require.Equal(t, 18, calcResult.spaceUsedByChildren)
	require.Equal(t, []int{6, 6, 6}, calcResult.actualSizes)
}

func TestMainAxisExtraSpaceWithEvenGrowth(t *testing.T) {
	minSizes := []int{
		6,
		3,
		3,
	}
	// Chose these numbers so growth happens evenly
	maxSizes := []int{
		12,
		6,
		6,
	}
	shouldGrow := []bool{
		true,
		true,
		true,
	}
	calcResult := calculateActualMainAxisSizes(minSizes, maxSizes, shouldGrow, 30)

	require.Equal(t, 30, calcResult.spaceUsedByChildren)
	require.Equal(t, []int{14, 8, 8}, calcResult.actualSizes)
}

func TestMainAxisExtraSpaceWithUnvenGrowth(t *testing.T) {
	minSizes := []int{
		5,
		3,
		3,
	}
	// Chose these numbers so growth happens evenly
	maxSizes := []int{
		10,
		5,
		5,
	}
	shouldGrow := []bool{
		false,
		true,
		false,
	}
	calcResult := calculateActualMainAxisSizes(minSizes, maxSizes, shouldGrow, 40)

	require.Equal(t, 40, calcResult.spaceUsedByChildren)
	require.Equal(t, []int{10, 25, 5}, calcResult.actualSizes)
}
