package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/box-layout-test/app_components/favorite_thing"
	"github.com/mieubrisse/box-layout-test/app_components/favorite_things_list"
	"github.com/mieubrisse/box-layout-test/components"
	"github.com/mieubrisse/box-layout-test/components/stylebox"
)

type App interface {
	components.InteractiveComponent
}

type appImpl struct {
	favoriteThingsList favorite_things_list.FavoriteThingsList

	root components.Component

	isFocused bool
}

func New() App {
	myFavoriteThings := []favorite_thing.FavoriteThing{
		favorite_thing.New().SetName("Pourover coffee").SetDescription("It takes so long to make though"),
		favorite_thing.New().SetName("Pizza").SetDescription("Pepperoni is the best"),
		favorite_thing.New().SetName("Jiu jitsu").SetDescription("Rolling all day"),
	}

	favoriteThingsList := favorite_things_list.New().SetThings(myFavoriteThings)

	root := stylebox.New(favoriteThingsList).SetStyle(lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()))
	return &appImpl{
		favoriteThingsList: favoriteThingsList,
		root:               root,
		isFocused:          false,
	}
}

func (a appImpl) GetContentMinMax() (minWidth, maxWidth, minHeight, maxHeight int) {
	return a.root.GetContentMinMax()
}

func (a appImpl) GetContentHeightForGivenWidth(width int) int {
	return a.root.GetContentHeightForGivenWidth(width)
}

func (a appImpl) View(width int, height int) string {
	return a.root.View(width, height)
}

func (a *appImpl) Update(msg tea.Msg) tea.Cmd {
	if !a.isFocused {
		return nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			things := a.favoriteThingsList.GetThings()
			thingNumber := len(things)
			newThing := favorite_thing.New().
				SetName(fmt.Sprintf("Thing #%v", thingNumber)).
				SetDescription(fmt.Sprintf("This is thing %v", thingNumber))
			things = append(things, newThing)
			a.favoriteThingsList.SetThings(things)
		} else if msg.String() == "backspace" {
			things := a.favoriteThingsList.GetThings()
			a.favoriteThingsList.SetThings(things[:len(things)-1])
		}
	}
	return nil
}

func (a *appImpl) SetFocus(isFocused bool) tea.Cmd {
	a.isFocused = true
	return nil
}

func (a appImpl) IsFocused() bool {
	return a.isFocused
}
