package flexbox

type FlexboxOpt func(Flexbox)

func WithDirection(direction Direction) FlexboxOpt {
	return func(box Flexbox) {
		box.SetDirection(direction)
	}
}

func WithHorizontalAlignment(alignment AxisAlignment) FlexboxOpt {
	return func(box Flexbox) {
		box.SetHorizontalAlignment(alignment)
	}
}

func WithVerticalAlignment(alignment AxisAlignment) FlexboxOpt {
	return func(box Flexbox) {
		box.SetVerticalAlignment(alignment)
	}
}
