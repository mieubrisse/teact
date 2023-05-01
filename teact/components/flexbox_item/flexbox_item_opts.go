package flexbox_item

// These are simply conveniences for the flexbox.NewWithContent , so that it's super easy to declare a single-item box
type FlexboxItemOpt func(item FlexboxItem)

func WithMinWidth(min FlexboxItemDimensionValue) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetMinWidth(min)
	}
}

func WithMaxWidth(max FlexboxItemDimensionValue) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetMaxWidth(max)
	}
}

func WithMinHeight(min FlexboxItemDimensionValue) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetMinHeight(min)
	}
}

func WithMaxHeight(max FlexboxItemDimensionValue) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetMaxHeight(max)
	}
}

func WithOverflowStyle(style OverflowStyle) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetOverflowStyle(style)
	}
}

func WithHorizontalGrowthFactor(growthFactor int) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetHorizontalGrowthFactor(growthFactor)
	}
}

func WithVerticalGrowthFactor(growthFactor int) FlexboxItemOpt {
	return func(item FlexboxItem) {
		item.SetVerticalGrowthFactor(growthFactor)
	}
}
