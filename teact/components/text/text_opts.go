package text

type TextOpt func(Text)

func WithContents(contents string) TextOpt {
	return func(text Text) {
		text.SetContents(contents)
	}
}

func WithAlign(align TextAlignment) TextOpt {
	return func(text Text) {
		text.SetTextAlignment(align)
	}
}
