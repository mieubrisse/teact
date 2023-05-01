package content_item

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/components"
	"github.com/mieubrisse/teact/components/flexbox"
	"github.com/mieubrisse/teact/components/flexbox_item"
	"github.com/mieubrisse/teact/components/stylebox"
	"github.com/mieubrisse/teact/components/text"
	"strings"
	"time"
)

var nameStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF4444"))

var tagsStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#4444FF"))

type ContentItem interface {
	components.Component

	GetTimestamp() time.Time

	GetName() string
	SetName(name string) ContentItem

	GetTags() []string
	SetTags(tags []string) ContentItem
}

type impl struct {
	timestamp time.Time
	name      string
	tags      []string

	nameText text.Text
	tagsText text.Text

	root components.Component
}

func New(timestamp time.Time, name string, tags []string) ContentItem {
	nameText := text.New(name)
	tagsText := text.New(strings.Join(tags, " "))

	root := flexbox.NewWithContents(
		flexbox_item.New(stylebox.New(nameText).SetStyle(nameStyle)).
			SetMinWidth(flexbox_item.FixedSize(20)).
			SetMaxWidth(flexbox_item.FixedSize(30)),
		flexbox_item.New(text.New(" ")).SetMinWidth(flexbox_item.FixedSize(1)),
		flexbox_item.New(stylebox.New(tagsText).SetStyle(tagsStyle)).SetHorizontalGrowthFactor(1),
	)

	return &impl{

		name:     name,
		tags:     tags,
		nameText: nameText,
		tagsText: tagsText,
		root:     root,
	}
}

func (f *impl) GetTimestamp() time.Time {
	return f.timestamp
}

func (f *impl) GetName() string {
	return f.name
}

func (f *impl) SetName(name string) ContentItem {
	f.name = name
	f.nameText.SetContents(name)
	return f
}

func (f *impl) GetTags() []string {
	return f.tags
}

func (f *impl) SetTags(tags []string) ContentItem {
	f.tags = tags
	f.tagsText.SetContents(strings.Join(tags, " "))
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
