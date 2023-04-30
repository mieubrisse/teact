package favorite_things_list

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/box-layout-test/app_components/favorite_thing"
	"github.com/mieubrisse/box-layout-test/components"
	"github.com/mieubrisse/box-layout-test/components/flexbox"
	"github.com/mieubrisse/box-layout-test/components/flexbox_item"
	"github.com/mieubrisse/box-layout-test/components/stylebox"
	"github.com/mieubrisse/box-layout-test/components/text"
)

type FavoriteThingsList interface {
	GetThings() []favorite_thing.FavoriteThing
	SetThings(elems []favorite_thing.FavoriteThing) FavoriteThingsList
}

type listImpl struct {
	entries []favorite_thing.FavoriteThing

	// Used to control the contents of the entry box
	entriesBox flexbox.Flexbox

	// The top element where everything will be rendered off of
	root components.Component
}

func New() FavoriteThingsList {
	titleText := text.New("My Favorite Things").SetTextAlignment(text.AlignCenter)
	titleComponent := stylebox.New(titleText).SetStyle(lipgloss.NewStyle().Bold(true))

	entriesBox := flexbox.New().
		SetDirection(flexbox.Column).
		SetHorizontalAlignment(flexbox.AlignStart).
		SetVerticalAlignment(flexbox.AlignCenter)
	styledEntryBox := stylebox.New(entriesBox).SetStyle(lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()))

	root := flexbox.NewWithContents(
		flexbox_item.New(titleComponent),
		flexbox_item.New(styledEntryBox),
	).SetDirection(flexbox.Column).
		SetVerticalAlignment(flexbox.AlignStart).
		SetHorizontalAlignment(flexbox.AlignCenter)

	return &listImpl{
		entries:    []favorite_thing.FavoriteThing{},
		root:       root,
		entriesBox: entriesBox,
	}
}

func (t listImpl) GetThings() []favorite_thing.FavoriteThing {
	return t.entries
}

func (t *listImpl) SetThings(entries []favorite_thing.FavoriteThing) FavoriteThingsList {
	t.entries = entries

	flexboxItems := make([]flexbox_item.FlexboxItem, len(entries))
	for idx, entry := range entries {
		flexboxItems[idx] = flexbox_item.New(entry)
	}

	t.entriesBox.SetChildren(flexboxItems)

	return t
}

func (t listImpl) GetContentMinMax() (minWidth, maxWidth, minHeight, maxHeight int) {
	return t.root.GetContentMinMax()
}

func (t listImpl) GetContentHeightForGivenWidth(width int) int {
	return t.root.GetContentHeightForGivenWidth(width)
}

func (t listImpl) View(width int, height int) string {
	return t.root.View(width, height)
}
