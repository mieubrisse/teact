package test_assertions

import (
	"github.com/mieubrisse/box-layout-test/components"
	"github.com/stretchr/testify/require"
	"testing"
)

// HeightAtWidthAssertion asserts that the component has a given height at a given width
type HeightAtWidthAssertion struct {
	Width          int
	ExpectedHeight int
}

func (assertion HeightAtWidthAssertion) Check(t *testing.T, component components.Component) {
	height := component.GetContentHeightForGivenWidth(assertion.Width)
	require.Equal(
		t,
		assertion.ExpectedHeight,
		height,
		"Expected the component to be height %v at width %v, but was %v",
		assertion.ExpectedHeight,
		assertion.Width,
		height,
	)
}

// Helper to create multiple height-at-width assertions
func GetHeightAtWidthAssertions(dimensions ...int) []ComponentAssertion {
	if len(dimensions)%2 != 0 {
		panic("Must provide dimensions in pairs of (width, height)")
	}

	result := make([]ComponentAssertion, 0, len(dimensions)/2)
	for i := 0; i < len(dimensions); i += 2 {
		result = append(
			result,
			HeightAtWidthAssertion{
				Width:          dimensions[i],
				ExpectedHeight: dimensions[i+1],
			},
		)
	}
	return result
}
