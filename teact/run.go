package teact

import (
	tea "github.com/charmbracelet/bubbletea"
	components2 "github.com/mieubrisse/teact/teact/components"
)

func RunTeactApp[T components2.Component](
	appComponent T,
	bubbleBathOptions []TeactOpt,
	teaOptions []tea.ProgramOption,
) (T, error) {
	model := New(appComponent, bubbleBathOptions...)

	finalModel, err := tea.NewProgram(model, teaOptions...).Run()
	castedModel := finalModel.(*impl)
	castedAppComponent := castedModel.app.(T)
	return castedAppComponent, err
}
