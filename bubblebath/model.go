package bubblebath

import (
	tea "github.com/charmbracelet/bubbletea"
	components2 "github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox"
	flexbox_item2 "github.com/mieubrisse/teact/teact/components/flexbox_item"
)

type BubbleBathOption func(*bubbleBathModel)

func WithInitCmd(cmd tea.Cmd) BubbleBathOption {
	return func(model *bubbleBathModel) {
		model.initCmd = cmd
	}
}

func WithQuitSequences(quitSequenceSet map[string]bool) BubbleBathOption {
	return func(model *bubbleBathModel) {
		model.quitSequenceSet = quitSequenceSet
	}
}

var defaultQuitSequenceSet = map[string]bool{
	"ctrl+c": true,
	"ctrl+d": true,
}

type bubbleBathModel struct {
	// The tea.Cmd that will be fired upon initialization
	initCmd tea.Cmd

	// Sequences matching String() of tea.KeyMsg that will quit the program
	quitSequenceSet map[string]bool

	appBox components2.Component

	app components2.Component

	width  int
	height int
}

// NewBubbleBathModel creates a new tea.Model for tea.NewProgram based off the given InteractiveComponent
func NewBubbleBathModel(app components2.Component, options ...BubbleBathOption) tea.Model {
	// We put the user's app in a box here so that we can get their app auto-resizing with the terminal
	appBox := flexbox.New().SetChildren([]flexbox_item2.FlexboxItem{
		flexbox_item2.New(app).
			// TODO allow these to be configured?
			SetMinWidth(flexbox_item2.MinContent).
			SetMaxWidth(flexbox_item2.MaxContent).
			SetHorizontalGrowthFactor(1).
			SetMinHeight(flexbox_item2.MinContent).
			SetMaxHeight(flexbox_item2.MaxContent).
			SetVerticalGrowthFactor(1),
	})
	result := &bubbleBathModel{
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

func (b bubbleBathModel) Init() tea.Cmd {
	return b.initCmd
}

func (b *bubbleBathModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	case components2.InteractiveComponent:
		cmd = app.Update(msg)
	}

	return b, cmd
}

func (b *bubbleBathModel) View() string {
	// We call these without using the results because:
	// 1) this is the three-phase cycle of our component rendering
	// 2) some components do caching of the phases, so to kick the cycle off we want to make sure we call them all
	b.appBox.GetContentMinMax()
	b.appBox.SetWidthAndGetDesiredHeight(b.width)
	return b.appBox.View(b.width, b.height)
}

/*
func (b bubbleBathModel) GetAppComponent() InteractiveComponent {
	return b.appComponent
}
*/

func RunBubbleBathProgram[T components2.Component](
	appComponent T,
	bubbleBathOptions []BubbleBathOption,
	teaOptions []tea.ProgramOption,
) (T, error) {
	model := NewBubbleBathModel(appComponent, bubbleBathOptions...)

	finalModel, err := tea.NewProgram(model, teaOptions...).Run()
	castedModel := finalModel.(*bubbleBathModel)
	castedAppComponent := castedModel.app.(T)
	return castedAppComponent, err
}
