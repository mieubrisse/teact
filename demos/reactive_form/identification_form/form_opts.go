package identification_form

type IdentificationFormOpts func(form IdentificationForm)

func WithFocus(isFocused bool) IdentificationFormOpts {
	return func(form IdentificationForm) {
		form.SetFocus(isFocused)
	}
}

func WithName(name string) IdentificationFormOpts {
	return func(form IdentificationForm) {
		form.SetName(name)
	}
}
