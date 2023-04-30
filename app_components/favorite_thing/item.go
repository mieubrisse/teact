package favorite_thing

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/box-layout-test/components"
	"github.com/mieubrisse/box-layout-test/components/flexbox"
	"github.com/mieubrisse/box-layout-test/components/flexbox_item"
	"github.com/mieubrisse/box-layout-test/components/stylebox"
	"github.com/mieubrisse/box-layout-test/components/text"
)

var nameStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF0000")).
	Bold(true)

var descriptionStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#0000FF")).
	Bold(true)

type FavoriteThing interface {
	components.Component

	GetName() string
	SetName(name string) FavoriteThing

	GetDescription() string
	SetDescription(desc string) FavoriteThing
}

type favoriteThingImpl struct {
	name        string
	description string

	nameText        text.Text
	descriptionText text.Text

	root components.Component
}

func New() FavoriteThing {
	nameText := text.New("")
	descriptionText := text.New("")

	root := flexbox.NewWithContents(
		flexbox_item.New(stylebox.New(nameText).SetStyle(nameStyle)),
		flexbox_item.New(stylebox.New(descriptionText).SetStyle(descriptionStyle)),
	).SetHorizontalAlignment(flexbox.AlignCenter)

	return &favoriteThingImpl{
		name:            "",
		description:     "",
		nameText:        nameText,
		descriptionText: descriptionText,
		root:            root,
	}
}

func (f *favoriteThingImpl) GetName() string {
	return f.name
}

func (f *favoriteThingImpl) SetName(name string) FavoriteThing {
	f.name = name
	f.nameText.SetContents(name)
	return f
}

func (f *favoriteThingImpl) GetDescription() string {
	return f.description
}

func (f *favoriteThingImpl) SetDescription(desc string) FavoriteThing {
	f.description = desc
	f.descriptionText.SetContents(desc)
	return f
}

func (f favoriteThingImpl) GetContentMinMax() (minWidth, maxWidth, minHeight, maxHeight int) {
	return f.root.GetContentMinMax()
}

func (f favoriteThingImpl) GetContentHeightForGivenWidth(width int) int {
	return f.root.GetContentHeightForGivenWidth(width)
}

func (f favoriteThingImpl) View(width int, height int) string {
	return f.root.View(width, height)
}
