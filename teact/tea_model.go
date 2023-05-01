package teact

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox"
	"github.com/mieubrisse/teact/teact/components/flexbox_item"
)

var defaultQuitSequenceSet = map[string]bool{
	"ctrl+c": true,
	"ctrl+d": true,
}

type TeactOpt func(*impl)

func WithInitCmd(cmd tea.Cmd) TeactOpt {
	return func(model *impl) {
		model.initCmd = cmd
	}
}

func WithQuitSequences(quitSequenceSet map[string]bool) TeactOpt {
	return func(model *impl) {
		model.quitSequenceSet = quitSequenceSet
	}
}

type TeactApp interface {
	// "Set" of quit sequences
	GetQuitSequences() map[string]bool
	SetQuitSequences(sequences map[string]bool) TeactApp
}

type impl struct {
	// The tea.Cmd that will be fired upon initialization
	initCmd tea.Cmd

	// Sequences matching String() of tea.KeyMsg that will quit the program
	quitSequenceSet map[string]bool

	appBox components.Component

	app components.Component

	width  int
	height int
}

// New creates a new tea.Model for tea.NewProgram running the given components.Component
func New(app components.Component, options ...TeactOpt) tea.Model {
	// We put the user's app in a box here so that we can get their app auto-resizing with the terminal
	appBox := flexbox.New(
		flexbox_item.New(
			app,
			flexbox_item.WithHorizontalGrowthFactor(1),
			flexbox_item.WithVerticalGrowthFactor(1),
		),
	)
	result := &impl{
		initCmd:         nil,
		quitSequenceSet: defaultQuitSequenceSet,
		appBox:          appBox,
		app:             app,
		width:           0,
		height:          0,
	}
	for _, opt := range options {
		opt(result)
	}
	return result
}

func (b impl) Init() tea.Cmd {
	return b.initCmd
}

func (b *impl) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if _, found := b.quitSequenceSet[msg.String()]; found {
			return b, tea.Quit

		}
	case tea.WindowSizeMsg:
		// b.appComponent.Resize(msg.Width, msg.Height)
		b.width = msg.Width
		b.height = msg.Height
		return b, nil
	}

	// Pass the message down to the app, if it's interactive
	var cmd tea.Cmd
	switch app := b.app.(type) {
	case components.InteractiveComponent:
		cmd = app.Update(msg)
	}

	return b, cmd
}

func (b *impl) View() string {
	// We call these without using the results because:
	// 1) this is the three-phase cycle of our component rendering
	// 2) some components do caching of the phases, so to kick the cycle off we want to make sure we call them all
	b.appBox.GetContentMinMax()
	b.appBox.SetWidthAndGetDesiredHeight(b.width)
	return b.appBox.View(b.width, b.height)
}
