package bio_card

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/demos/secret_agent_terminal/colors"
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox"
	"github.com/mieubrisse/teact/teact/components/flexbox_item"
	"github.com/mieubrisse/teact/teact/components/stylebox"
	"github.com/mieubrisse/teact/teact/components/text"
	"github.com/mieubrisse/teact/teact/utilities"
	"strconv"
)

const (
	unknownName = "Anonymous Agent"
)

var normalTextStyle = utilities.NewStyle(
	utilities.WithForeground(colors.Platinum),
)
var nameStyle = utilities.NewStyle(
	utilities.WithForeground(colors.Tomato),
	utilities.WithBold(true),
)
var ageStyle = utilities.NewStyle(
	utilities.WithForeground(colors.VividSkyBlue),
	utilities.WithBold(true),
)

type bioCardImpl struct {
	components.Component

	name string
	age  int

	row flexbox.Flexbox
}

func New() BioCard {
	row := flexbox.NewWithOpts(
		[]flexbox_item.FlexboxItem{},
		flexbox.WithHorizontalAlignment(flexbox.AlignCenter),
		flexbox.WithVerticalAlignment(flexbox.AlignCenter),
	)
	result := &bioCardImpl{
		Component: row,
		name:      "",
		age:       0,
		row:       row,
	}
	result.updateFlexbox()
	return result
}

func (impl *bioCardImpl) SetName(name string) BioCard {
	impl.name = name
	impl.updateFlexbox()
	return impl
}

func (impl *bioCardImpl) SetAge(age int) BioCard {
	impl.age = age
	impl.updateFlexbox()
	return impl
}

func (impl *bioCardImpl) updateFlexbox() {
	name := impl.name
	if name == "" {
		name = unknownName
	}

	texts := []string{
		"Hello, ",
		name,
		". ",
	}
	styles := []lipgloss.Style{
		normalTextStyle,
		nameStyle,
		normalTextStyle,
	}

	// TODO we reallyyyyy need an inline element

	if impl.age == 0 {
		texts = append(texts, "We don't know how old you are.")
		styles = append(styles, normalTextStyle)
	} else {
		texts = append(texts,
			"You are ",
			strconv.Itoa(impl.age),
			" years old.",
		)
		styles = append(styles,
			normalTextStyle,
			ageStyle,
			normalTextStyle,
		)
	}

	flexboxItems := make([]flexbox_item.FlexboxItem, len(texts))
	for idx, textFragment := range texts {
		fragmentStyle := styles[idx]
		flexboxItems[idx] = flexbox_item.New(
			stylebox.New(
				text.New(textFragment),
				stylebox.WithExistingStyle(fragmentStyle),
			),
		)
	}

	impl.row.SetChildren(flexboxItems)
}
