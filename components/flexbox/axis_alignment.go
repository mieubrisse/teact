package flexbox

// The percentage from the start that alignment should be done
// See lipgloss.Position for more
type AxisAlignment float64

const (
	// Elements will be at the start of the flexbox (as determined by the Direction)
	// Corresponds to "flex-justify: flex-start"
	AlignStart AxisAlignment = 0.0

	// NOTE: in order to see this in effect, you must have
	// Corresponds to "flex-justify: center"
	AlignCenter = 0.5

	// Elements will be pushed to the end of the flexbox (as determined by the Direction)
	// Corresponds to "flex-justify: flex-end"
	AlignEnd = 1.0
)
