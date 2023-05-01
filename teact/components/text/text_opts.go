package text

type TextOpt func(Text)

func WithAlign(align TextAlignment) TextOpt {
	return func(text Text) {
		text.SetTextAlignment(align)
	}
}
