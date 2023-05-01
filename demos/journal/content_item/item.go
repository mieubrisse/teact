package content_item

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/components"
	"github.com/mieubrisse/teact/components/flexbox"
	"github.com/mieubrisse/teact/components/flexbox_item"
	"github.com/mieubrisse/teact/components/highlightable_list"
	"github.com/mieubrisse/teact/components/stylebox"
	"github.com/mieubrisse/teact/components/text"
	"strings"
	"time"
)

var highlightedBackgroundColor = lipgloss.Color("#333333")

var nameStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF")).
	Padding(0, 1, 0, 0)

var tagsStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF5555"))

type ContentItem interface {
	components.Component
	highlightable_list.HighlightableComponent

	GetTimestamp() time.Time

	GetName() string
	SetName(name string) ContentItem

	GetTags() []string
	SetTags(tags []string) ContentItem
}

type impl struct {
	// Root item
	components.Component

	timestamp time.Time
	name      string
	tags      []string

	nameText text.Text
	tagsText text.Text

	isHighlighted             bool
	toChangeOnHighlightToggle []stylebox.Stylebox
}

func New(timestamp time.Time, name string, tags []string) ContentItem {
	nameText := text.New(name)
	tagsText := text.New(strings.Join(tags, " "))

	styledName := stylebox.New(nameText).SetStyle(nameStyle)
	styledTags := stylebox.New(nameText).SetStyle(tagsStyle)

	itemsRow := flexbox.NewWithContents(
		flexbox_item.New(styledName).
			SetMinWidth(flexbox_item.FixedSize(20)).
			SetMaxWidth(flexbox_item.FixedSize(30)),
		flexbox_item.New(styledTags).SetHorizontalGrowthFactor(1),
	)

	styledItemsRow := stylebox.New(itemsRow)

	toChangeOnHighlightToggle := []stylebox.Stylebox{
		styledName,
		styledTags,
		styledItemsRow,
	}

	return &impl{
		Component:                 styledItemsRow,
		timestamp:                 timestamp,
		name:                      name,
		tags:                      tags,
		nameText:                  nameText,
		tagsText:                  tagsText,
		isHighlighted:             false,
		toChangeOnHighlightToggle: toChangeOnHighlightToggle,
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

func (f *impl) IsHighlighted() bool {
	return f.isHighlighted
}

func (f *impl) SetHighlight(isHighlighted bool) highlightable_list.HighlightableComponent {
	f.isHighlighted = isHighlighted

	for _, box := range f.toChangeOnHighlightToggle {
		if isHighlighted {
			box.GetStyle().Bold(true).Background(highlightedBackgroundColor)
		} else {
			box.GetStyle().UnsetBold().UnsetBackground()
		}
	}

	return f
}
