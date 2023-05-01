package text_input

type TextInputOpt func(input TextInput)

func WithFocus(isFocused bool) TextInputOpt {
	return func(input TextInput) {
		input.SetFocus(isFocused)
	}
}

func WithValue(value string) TextInputOpt {
	return func(input TextInput) {
		input.SetValue(value)
	}
}
