package utilities

import "github.com/charmbracelet/lipgloss"

// Procrustean coersion to make the given string exactly fit the dimensions we want
func Coerce(str string, actualWidth int, actualHeight int) string {
	// Truncate, in case we're too long
	truncated := lipgloss.NewStyle().MaxWidth(actualWidth).MaxHeight(actualHeight).Render(str)

	// Expand to fill (in case we're small)
	return lipgloss.NewStyle().Width(actualWidth).Height(actualHeight).Render(truncated)
}
