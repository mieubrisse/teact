package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/demos/journal/content_item"
	components2 "github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/highlightable_list"
	"github.com/mieubrisse/teact/teact/components/stylebox"
	"time"
)

type App interface {
	components2.InteractiveComponent
}

type appImpl struct {
	components2.Component

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
	itemsList.SetFocus(true)

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

	return a.itemsList.Update(msg)
}

func (a *appImpl) SetFocus(isFocused bool) tea.Cmd {
	a.isFocused = true
	return nil
}

func (a appImpl) IsFocused() bool {
	return a.isFocused
}
