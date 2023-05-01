package teact

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/flexbox"
	"github.com/mieubrisse/teact/teact/components/flexbox_item"
)

// The root tea.Model that runs the Teact framework
type TeactModel interface {
	tea.Model

	// The user's component being controlled as the root of this Teact app
	GetInnerComponent() components.Component

	// "Set" of quit sequences
	GetQuitSequences() map[string]bool
	SetQuitSequences(sequences map[string]bool) TeactModel
}

var defaultQuitSequenceSet = map[string]bool{
	"ctrl+c": true,
	"ctrl+d": true,
}

type teactAppModelImpl struct {
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
func New[T components.Component](app T, options ...TeactModelOpt) TeactModel {
	// We put the user's app in a box here so that we can get their app auto-resizing with the terminal
	appBox := flexbox.New(
		flexbox_item.New(
			app,
			flexbox_item.WithHorizontalGrowthFactor(1),
			flexbox_item.WithVerticalGrowthFactor(1),
		),
	)
	result := &teactAppModelImpl{
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

func (impl teactAppModelImpl) Init() tea.Cmd {
	return impl.initCmd
}

func (impl *teactAppModelImpl) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if _, found := impl.quitSequenceSet[msg.String()]; found {
			return impl, tea.Quit

		}
	case tea.WindowSizeMsg:
		impl.width = msg.Width
		impl.height = msg.Height
		return impl, nil
	}

	// Pass the message down to the app, if it's interactive
	var cmd tea.Cmd
	switch app := impl.app.(type) {
	case components.InteractiveComponent:
		cmd = app.Update(msg)
	}

	return impl, cmd
}

func (impl *teactAppModelImpl) View() string {
	// We call these without using the results because:
	// 1) this is the three-phase cycle of our component rendering
	// 2) some components do caching of the phases, so to kick the cycle off we want to make sure we call them all
	impl.appBox.GetContentMinMax()
	impl.appBox.SetWidthAndGetDesiredHeight(impl.width)
	return impl.appBox.View(impl.width, impl.height)
}

func (impl teactAppModelImpl) GetQuitSequences() map[string]bool {
	return impl.quitSequenceSet
}

func (impl *teactAppModelImpl) SetQuitSequences(sequences map[string]bool) TeactModel {
	impl.quitSequenceSet = sequences
	return impl
}

func (impl teactAppModelImpl) GetInnerComponent() components.Component {
	return impl.app
}
