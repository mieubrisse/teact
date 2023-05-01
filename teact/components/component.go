package components

type Component interface {
	// This is used during the X-expansion phase, where each child "expands" its min and max widths up to its parent
	// During this stage, each element is growing in the X direction; there is no concept of a viewport
	GetContentMinMax() (minWidth, maxWidth, minHeight, maxHeight int)

	// This is used during the Y-expansion phase, where a viewport width is known and now we're determining heights
	// This method should do any necessary reflowing, and then get the desired height
	// If you want to do any reflowing based on the actual size of the viewport (e.g. maybe stacking a sidebar vertically for small viewports), this is the place to do it
	SetWidthAndGetDesiredHeight(actualWidth int) int

	// The 'width' will be the same width that was passed in to GetContentHeightForGivenWidth, allowing for some caching
	// of calculation results between the two
	// TODO maybe return Optional[string], so that we can indicate "there is no content at all"?
	View(actualWidth int, actualHeight int) string
}
