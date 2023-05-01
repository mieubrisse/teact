package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/app_components/content_item"
	"github.com/mieubrisse/teact/components"
	"github.com/mieubrisse/teact/components/highlightable_list"
	"github.com/mieubrisse/teact/components/stylebox"
	"time"
)

type App interface {
	components.InteractiveComponent
}

type appImpl struct {
	components.Component

	itemsList highlightable_list.HighlightableList[content_item.ContentItem]

	isFocused bool
}

func New() App {
	items := []content_item.ContentItem{
		content_item.New(time.Now(), "foo.md", []string{"general-reference"}),
		content_item.New(time.Now(), "bar-bang-baz.md", []string{"project-support/starlark"}),
		content_item.New(time.Now(), "something-else.md", []string{"general-reference/wealthdraft"}),
	}

	itemsList := highlightable_list.New[content_item.ContentItem]()
	itemsList.SetItems(items)
	itemsList.SetHighlightedIdx(0)

	root := stylebox.New(itemsList).SetStyle(lipgloss.NewStyle().Padding(1, 2))
	root = stylebox.New(root).SetStyle(lipgloss.NewStyle().Padding(1, 2))
	return &appImpl{
		Component: root,
		itemsList: itemsList,
		isFocused: false,
	}
}

func (a *appImpl) Update(msg tea.Msg) tea.Cmd {
	if !a.isFocused {
		return nil
	}

	/*
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				things := a.contentList.GetItems()
				thingNumber := len(things)
				newThing := content_item.New().
					SetName(fmt.Sprintf("Thing #%v", thingNumber)).
					SetTags(fmt.Sprintf("This is thing %v", thingNumber))
				things = append(things, newThing)
				a.contentList.SetItems(things)
			} else if msg.String() == "backspace" {
				things := a.contentList.GetItems()
				a.contentList.SetItems(things[:len(things)-1])
			}
		}

	*/
	return nil
}

func (a *appImpl) SetFocus(isFocused bool) tea.Cmd {
	a.isFocused = true
	return nil
}

func (a appImpl) IsFocused() bool {
	return a.isFocused
}
