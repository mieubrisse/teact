package keypress_counter

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mieubrisse/teact/teact/components"
	"github.com/mieubrisse/teact/teact/components/text"
	"github.com/mieubrisse/teact/teact/utilities"
)

type KeypressCounter interface {
	components.InteractiveComponent
}

type keypressCounterImpl struct {
	components.Component

	keysPressed int
	output      text.Text
}

func New() KeypressCounter {
	output := text.New("")
	result := &keypressCounterImpl{
		Component:   output,
		keysPressed: 0,
		output:      output,
	}
	result.updateOutputText()
	return result
}

func (k *keypressCounterImpl) Update(msg tea.Msg) tea.Cmd {
	if utilities.GetMaybeKeyMsgStr(msg) != "" {
		k.keysPressed += 1
		k.updateOutputText()
	}
	return nil
}

func (k keypressCounterImpl) SetFocus(isFocused bool) tea.Cmd {
	return nil
}

func (k keypressCounterImpl) IsFocused() bool {
	return true
}

func (b *keypressCounterImpl) updateOutputText() {
	b.output.SetContents(fmt.Sprintf("You've pressed %v keys", b.keysPressed))
}
