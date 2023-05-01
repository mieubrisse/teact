package teact

import tea "github.com/charmbracelet/bubbletea"

type TeactModelOpt func(*teactAppModelImpl)

func WithInitCmd(cmd tea.Cmd) TeactModelOpt {
	return func(model *teactAppModelImpl) {
		model.initCmd = cmd
	}
}

func WithQuitSequences(quitSequenceSet map[string]bool) TeactModelOpt {
	return func(model *teactAppModelImpl) {
		model.quitSequenceSet = quitSequenceSet
	}
}
