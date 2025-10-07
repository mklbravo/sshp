package tui

import "github.com/charmbracelet/lipgloss"

type stylesStruct struct {
	mauve    lipgloss.Style
	sapphire lipgloss.Style
	sky      lipgloss.Style
	teal     lipgloss.Style
}

var styles = stylesStruct{
	mauve:    lipgloss.NewStyle().Foreground(lipgloss.Color("#cba6f7")),
	sapphire: lipgloss.NewStyle().Foreground(lipgloss.Color("#74c7ec")),
	sky:      lipgloss.NewStyle().Foreground(lipgloss.Color("#89dceb")),
	teal:     lipgloss.NewStyle().Foreground(lipgloss.Color("#94e2d5")),
}
