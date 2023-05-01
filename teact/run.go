package teact

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/teact/components"
)

// Simple way to run a Teact program
// If you need more complex configuration, use RunTeactFromModel
func Run[T components.Component](
	yourApp T,
	bubbleTeaOpts ...tea.ProgramOption,
) (T, error) {
	model := New(yourApp)
	finalModel, err := tea.NewProgram(model, bubbleTeaOpts...).Run()
	castedModel := finalModel.(TeactModel)
	castedUserComponent := castedModel.GetInnerComponent().(T)
	return castedUserComponent, err
}
