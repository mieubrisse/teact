package test_assertions

import (
	"github.com/mieubrisse/box-layout-test/components"
	"testing"
)

type ComponentAssertion interface {
	Check(t *testing.T, component components.Component)
}

// Helper to make working with groups of assertions easier
func FlattenAssertionGroups(assertionGroups ...[]ComponentAssertion) []ComponentAssertion {
	numAssertions := 0
	for _, group := range assertionGroups {
		numAssertions += len(group)
	}

	result := make([]ComponentAssertion, 0, numAssertions)
	for _, group := range assertionGroups {
		result = append(result, group...)
	}
	return result
}

// Run a group of assertions against the component
func CheckAll(t *testing.T, assertionGroup []ComponentAssertion, component components.Component) {
	for _, assertion := range assertionGroup {
		assertion.Check(t, component)
	}
}

func GetDefaultAssertions() []ComponentAssertion {
	return FlattenAssertionGroups(
		// Every component should be zero height when zero width
		GetHeightAtWidthAssertions(0, 0),

		// A zero height or width should always in an empty string
		GetRenderedContentAssertion(1, 0, ""),
		GetRenderedContentAssertion(0, 1, ""),
		GetRenderedContentAssertion(0, 0, ""),
	)
}
