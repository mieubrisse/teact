package grid

import (
	"github.com/mieubrisse/teact/components"
	"github.com/mieubrisse/teact/utilities"
)

type Grid interface {
	components.Component
}

type gridImpl struct {
	// TODO make this more genericized
	templateColumnFractions []int

	components []components.Component

	// TODO add grid lines
}

func New(templateColumnFractions []int, components []components.Component) Grid {
	return &gridImpl{
		templateColumnFractions: nil,
	}
}

func (g gridImpl) GetContentMinMax() (minWidth, maxWidth, minHeight, maxHeight int) {
	width := 0
	for _, value := range g.templateColumnFractions {
		width += value
	}

	return width, width, 10, 10
}

func (g gridImpl) GetContentHeightForGivenWidth(width int) int {
	return 10
}

func (g gridImpl) View(width int, height int) string {
	sizes := make([]int, len(g.templateColumnFractions))
	sizes = utilities.DistributeSpaceByWeight(width, sizes, g.templateColumnFractions)
}
