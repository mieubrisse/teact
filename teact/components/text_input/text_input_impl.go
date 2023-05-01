package text_input

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mieubrisse/teact/teact/utilities"
	"github.com/muesli/reflow/wordwrap"
	"strings"
)

type textInputImpl struct {
	innerInput textinput.Model

	isFocused bool
}

func New(opts ...TextInputOpt) TextInput {
	innerInput := textinput.New()
	result := &textInputImpl{
		innerInput: innerInput,
		isFocused:  false,
	}
	for _, opt := range opts {
		opt(result)
	}
	return result
}

func (i textInputImpl) GetContentMinMax() (int, int, int, int) {
	value := i.innerInput.Value()

	maxWidth := lipgloss.Width(value)
	minHeight := lipgloss.Height(value)

	minWidth := 0
	for _, field := range strings.Fields(value) {
		minWidth = utilities.GetMaxInt(minWidth, len(field))
	}

	reflowed := wordwrap.String(value, minWidth)
	maxHeight := lipgloss.Height(reflowed)

	return minWidth, maxWidth, minHeight, maxHeight
}

func (i *textInputImpl) SetWidthAndGetDesiredHeight(actualWidth int) int {
	value := i.innerInput.Value()
	reflowed := wordwrap.String(value, actualWidth)
	return lipgloss.Height(reflowed)
}

func (i *textInputImpl) View(actualWidth int, actualHeight int) string {
	value := i.innerInput.Value()
	reflowed := wordwrap.String(value, actualWidth)

	return utilities.Coerce(reflowed, actualWidth, actualHeight)
}

func (i *textInputImpl) Update(msg tea.Msg) tea.Cmd {
	if !i.isFocused {
		return nil
	}

	var cmd tea.Cmd
	i.innerInput, cmd = i.innerInput.Update(msg)
	return cmd
}

func (i *textInputImpl) SetFocus(isFocused bool) tea.Cmd {
	i.isFocused = isFocused
	if isFocused {
		i.innerInput.Focus()
	} else {
		i.innerInput.Blur()
	}
	return nil
}

func (i *textInputImpl) IsFocused() bool {
	return i.isFocused
}

func (i *textInputImpl) GetValue() string {
	return i.innerInput.Value()
}

func (i *textInputImpl) SetValue(value string) TextInput {
	i.innerInput.SetValue(value)
	return i
}
