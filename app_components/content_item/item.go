package content_item

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/components"
	"github.com/mieubrisse/teact/components/flexbox"
	"github.com/mieubrisse/teact/components/flexbox_item"
	"github.com/mieubrisse/teact/components/stylebox"
	"github.com/mieubrisse/teact/components/text"
)

var nameStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF4444")).
	Border(lipgloss.NormalBorder())

var tagsStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#4444FF")).
	Border(lipgloss.NormalBorder())

type ContentItem interface {
	components.Component

	GetName() string
	SetName(name string) ContentItem

	GetDescription() string
	SetDescription(desc string) ContentItem
}

type impl struct {
	name        string
	description string

	nameText        text.Text
	descriptionText text.Text

	root components.Component
}

func New() ContentItem {
	nameText := text.New("")
	descriptionText := text.New("")

	root := flexbox.NewWithContents(
		flexbox_item.New(stylebox.New(nameText).SetStyle(nameStyle)).
			SetMinWidth(flexbox_item.FixedSize(40)).
			SetMaxWidth(flexbox_item.MaxAvailable),
		flexbox_item.New(text.New(" ")).SetMinWidth(flexbox_item.FixedSize(1)),
		flexbox_item.New(stylebox.New(descriptionText).SetStyle(tagsStyle)),
	)

	return &impl{
		name:            "",
		description:     "",
		nameText:        nameText,
		descriptionText: descriptionText,
		root:            root,
	}
}

func (f *impl) GetName() string {
	return f.name
}

func (f *impl) SetName(name string) ContentItem {
	f.name = name
	f.nameText.SetContents(name)
	return f
}

func (f *impl) GetDescription() string {
	return f.description
}

func (f *impl) SetDescription(desc string) ContentItem {
	f.description = desc
	f.descriptionText.SetContents(desc)
	return f
}

func (f impl) GetContentMinMax() (minWidth, maxWidth, minHeight, maxHeight int) {
	return f.root.GetContentMinMax()
}

func (f impl) GetContentHeightForGivenWidth(width int) int {
	return f.root.GetContentHeightForGivenWidth(width)
}

func (f impl) View(width int, height int) string {
	return f.root.View(width, height)
}
