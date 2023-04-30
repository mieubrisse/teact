package components

type Component interface {
	// This is used during the X-expansion phase, where each child "expands" its min and max widths up to its parent
	// During this stage, each element is growing in the X direction; there is no concept of a viewport
	GetContentMinMax() (minWidth, maxWidth, minHeight, maxHeight int)

	// This is used during the Y-expansion phase, where a viewport width is known and now we're expanding vertically
	// (which requires us knowing the heights of elements so we can size appropriately)
	// It will give the component's desired height for the given width
	GetContentHeightForGivenWidth(width int) int

	// The 'width' will be the same width that was passed in to GetContentHeightForGivenWidth, allowing for some caching
	// of calculation results between the two
	// TODO maybe return Optional[string], so that we can indicate "there is no content at all"?
	View(width int, height int) string
}
