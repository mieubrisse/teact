package style

import "github.com/charmbracelet/lipgloss"

// Contains wrappers around Lipgloss' style, to turn Lipgloss' long, fluent methods into more concise ones

type StyleOpt func(style lipgloss.Style) lipgloss.Style

func WithForeground(color lipgloss.Color) StyleOpt {
	return func(style lipgloss.Style) lipgloss.Style {
		return style.Foreground(color)
	}
}

func WithBackground(color lipgloss.Color) StyleOpt {
	return func(style lipgloss.Style) lipgloss.Style {
		return style.Background(color)
	}
}

func WithBold(isBold bool) StyleOpt {
	return func(style lipgloss.Style) lipgloss.Style {
		return style.Bold(isBold)
	}
}

func WithItalic(isItalic bool) StyleOpt {
	return func(style lipgloss.Style) lipgloss.Style {
		return style.Italic(isItalic)
	}
}

func WithUnderline(isUnderline bool) StyleOpt {
	return func(style lipgloss.Style) lipgloss.Style {
		return style.Underline(isUnderline)
	}
}

func WithUnderlineSpaces(isUnderlineSpaces bool) StyleOpt {
	return func(style lipgloss.Style) lipgloss.Style {
		return style.UnderlineSpaces(isUnderlineSpaces)
	}
}

func WithStrikethrough(isStrikethrough bool) StyleOpt {
	return func(style lipgloss.Style) lipgloss.Style {
		return style.Strikethrough(isStrikethrough)
	}
}

func WithStrikethroughSpaces(isStrikethroughSpaces bool) StyleOpt {
	return func(style lipgloss.Style) lipgloss.Style {
		return style.StrikethroughSpaces(isStrikethroughSpaces)
	}
}

func WithFaint(isFaint bool) StyleOpt {
	return func(style lipgloss.Style) lipgloss.Style {
		return style.Faint(isFaint)
	}
}

func WithBlink(isBlink bool) StyleOpt {
	return func(style lipgloss.Style) lipgloss.Style {
		return style.Blink(isBlink)
	}
}

func WithBorder(border lipgloss.Border, sides ...bool) StyleOpt {
	return func(style lipgloss.Style) lipgloss.Style {
		return style.Border(border, sides...)
	}
}

func WithPadding(padding ...int) StyleOpt {
	return func(style lipgloss.Style) lipgloss.Style {
		return style.Padding(padding...)
	}
}

func NewStyle(opts ...StyleOpt) lipgloss.Style {
	result := lipgloss.NewStyle()
	for _, opt := range opts {
		result = opt(result)
	}
	return result
}

// TODO from existing style
