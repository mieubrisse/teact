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
	desiredSizes := []int{
		10,
		5,
		7,
	}
	shouldGrow := []bool{
		false,
		false,
		false,
	}
	calcResult := calculateActualCrossAxisSizes(desiredSizes, shouldGrow, 6)

	require.Equal(t, calcResult.spaceUsedByChildren, 6)
	require.Equal(t, calcResult.actualSizes, []int{6, 5, 6})
}

func TestCrossAxisTruncationWithGrowth(t *testing.T) {
	desiredSizes := []int{
		10,
		5,
		7,
	}
	shouldGrow := []bool{
		true,
		true,
		true,
	}
	calcResult := calculateActualCrossAxisSizes(desiredSizes, shouldGrow, 6)

	require.Equal(t, calcResult.spaceUsedByChildren, 6)
	require.Equal(t, calcResult.actualSizes, []int{6, 6, 6})
}

func TestCrossAxisExtraSpaceNoGrowth(t *testing.T) {
	desiredSizes := []int{
		10,
		5,
		7,
	}
	shouldGrow := []bool{
		false,
		false,
		false,
	}
	calcResult := calculateActualCrossAxisSizes(desiredSizes, shouldGrow, 12)

	require.Equal(t, calcResult.spaceUsedByChildren, 10)
	require.Equal(t, calcResult.actualSizes, []int{10, 5, 7})
}

func TestCrossAxisExtraSpaceWithGrowth(t *testing.T) {
	desiredSizes := []int{
		10,
		5,
		7,
	}
	shouldGrow := []bool{
		true,
		true,
		true,
	}
	calcResult := calculateActualCrossAxisSizes(desiredSizes, shouldGrow, 12)

	require.Equal(t, calcResult.spaceUsedByChildren, 12)
	require.Equal(t, calcResult.actualSizes, []int{12, 12, 12})
}

// ====================================================================================================
//
//	Main Axis Tests
//
// ====================================================================================================
func TestMainAxisExtraSpaceNoGrowth(t *testing.T) {
	desiredSizes := []int{
		10,
		5,
		7,
	}
	shouldGrow := []bool{
		false,
		false,
		false,
	}
	calcResult := calculateActualMainAxisSizes(desiredSizes, shouldGrow, 30)

	require.Equal(t, calcResult.spaceUsedByChildren, 22)
	require.Equal(t, calcResult.actualSizes, []int{10, 5, 7})
}

func TestMainAxisExtraSpaceWithEvenGrowth(t *testing.T) {
	// Chose these numbers so growth happens evenly
	desiredSizes := []int{
		10,
		5,
		5,
	}
	shouldGrow := []bool{
		true,
		true,
		true,
	}
	calcResult := calculateActualMainAxisSizes(desiredSizes, shouldGrow, 40)

	require.Equal(t, calcResult.spaceUsedByChildren, 40)
	require.Equal(t, calcResult.actualSizes, []int{20, 10, 10})
}

func TestMainAxisExtraSpaceWithUnvenGrowth(t *testing.T) {
	// Chose these numbers so growth happens evenly
	desiredSizes := []int{
		10,
		5,
		5,
	}
	shouldGrow := []bool{
		false,
		true,
		false,
	}
	calcResult := calculateActualMainAxisSizes(desiredSizes, shouldGrow, 40)

	require.Equal(t, calcResult.spaceUsedByChildren, 40)
	require.Equal(t, calcResult.actualSizes, []int{10, 25, 5})
}

func TestMainAxisShrink(t *testing.T) {
	desiredSizes := []int{
		8,
		2,
		2,
	}
	shouldGrow := []bool{
		false,
		false,
		false,
	}
	calcResult := calculateActualMainAxisSizes(desiredSizes, shouldGrow, 6)

	require.Equal(t, calcResult.spaceUsedByChildren, 6)
	require.Equal(t, calcResult.actualSizes, []int{4, 1, 1})
}
