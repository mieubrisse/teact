package flexbox_item

type FlexboxItemDimensionValue interface {
	// Given a min and a max, gets the corresponding size based on what FlexboxItemDimensionValue this is
	getSizeRetriever() func(min, max int) int
}

// Indicates a size == the minimum content size of the item, which:
// - For width is the size of the item if all wrapping opportunities are taken (basically, the length of the longest word)
// - For height is the height of the item when no word-wrapping is done
var MinContent = &dimensionValueImpl{
	sizeRetriever: func(min, max int) int {
		return min
	},
}

// Indicates a size == the maximum content of the item, which is the size of the item without any wrapping applied
// - For width, this is basically, the length of the longest line
// - For height, this is the height of the item when the maximum possible word-wrapping is done
var MaxContent = &dimensionValueImpl{
	sizeRetriever: func(min, max int) int {
		return max
	},
}

// Indicates a fixed size
func FixedSize(size int) FlexboxItemDimensionValue {
	return &dimensionValueImpl{
		sizeRetriever: func(min, max int) int {
			return size
		},
	}
}

// ====================================================================================================
//
//	Private
//
// ====================================================================================================
// This type represents values for a flexbox item dimension (height or width)
type dimensionValueImpl struct {
	// Given a min and a max, gets the corresponding size based on what FlexboxItemDimensionValue this is
	sizeRetriever func(min, max int) int
}

func (impl dimensionValueImpl) getSizeRetriever() func(min int, max int) int {
	return impl.sizeRetriever
}
