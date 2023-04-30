package test_assertions

import (
	"github.com/mieubrisse/box-layout-test/components"
	"github.com/stretchr/testify/require"
	"testing"
)

// RenderedContentAssertion asserts that the component renders the given output
type RenderedContentAssertion struct {
	Width           int
	Height          int
	ExpectedContent string
}

func (v RenderedContentAssertion) Check(t *testing.T, component components.Component) {
	output := component.View(v.Width, v.Height)
	require.Equal(t, v.ExpectedContent, output)
}

// This returns an array to make it very easy to slot into FlattenAssertionGroups
func GetRenderedContentAssertion(width int, height int, expectedContent string) []ComponentAssertion {
	return []ComponentAssertion{
		RenderedContentAssertion{
			Width:           width,
			Height:          height,
			ExpectedContent: expectedContent,
		},
	}
}
