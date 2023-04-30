package test_assertions

import (
	"github.com/mieubrisse/box-layout-test/components"
	"github.com/stretchr/testify/require"
	"testing"
)

// ContentSizeAssertion asserts the min/max width/height for the component
type ContentSizeAssertion struct {
	ExpectedMinWidth  int
	ExpectedMaxWidth  int
	ExpectedMinHeight int
	ExpectedMaxHeight int
}

func (assertion ContentSizeAssertion) Check(t *testing.T, component components.Component) {
	minWidth, maxWidth, minHeight, maxHeight := component.GetContentMinMax()
	require.Equal(
		t,
		assertion.ExpectedMinWidth,
		minWidth,
		"Expected the component's minWidth to be %v but was %v",
		assertion.ExpectedMinWidth,
		minWidth,
	)
	require.Equal(
		t,
		assertion.ExpectedMaxWidth,
		maxWidth,
		"Expected the component's maxWidth to be %v but was %v",
		assertion.ExpectedMaxWidth,
		maxWidth,
	)
	require.Equal(
		t,
		assertion.ExpectedMinHeight,
		minHeight,
		"Expected the component's minHeight to be %v but was %v",
		assertion.ExpectedMinHeight,
		minHeight,
	)
	require.Equal(
		t,
		assertion.ExpectedMaxHeight,
		maxHeight,
		"Expected the component's maxHeight to be %v but was %v",
		assertion.ExpectedMaxHeight,
		maxHeight,
	)
}

// Helper to create multiple content size assertions
func GetContentSizeAssertions(dimensions ...int) []ComponentAssertion {
	if len(dimensions)%4 != 0 {
		panic("Must provide dimensions in pairs of (minWidth, maxWidth, minHeight, maxHeight)")
	}

	result := make([]ComponentAssertion, 0, len(dimensions)/2)
	for i := 0; i < len(dimensions); i += 4 {
		result = append(
			result,
			ContentSizeAssertion{
				ExpectedMinWidth:  dimensions[i],
				ExpectedMaxWidth:  dimensions[i+1],
				ExpectedMinHeight: dimensions[i+2],
				ExpectedMaxHeight: dimensions[i+3],
			},
		)
	}
	return result
}
