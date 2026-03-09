package tui

import "github.com/charmbracelet/lipgloss"

type catppuccinMochaColorStyles struct {
	blue      lipgloss.Style
	flamingo  lipgloss.Style
	green     lipgloss.Style
	maroon    lipgloss.Style
	mauve     lipgloss.Style
	peach     lipgloss.Style
	pink      lipgloss.Style
	red       lipgloss.Style
	lavender  lipgloss.Style
	rosewater lipgloss.Style
	sapphire  lipgloss.Style
	sky       lipgloss.Style
	teal      lipgloss.Style
	yellow    lipgloss.Style
}

var colorStyle = catppuccinMochaColorStyles{
	blue:      lipgloss.NewStyle().Foreground(lipgloss.Color("#89b4fa")),
	flamingo:  lipgloss.NewStyle().Foreground(lipgloss.Color("#f2cdcd")),
	green:     lipgloss.NewStyle().Foreground(lipgloss.Color("#a6e3a1")),
	maroon:    lipgloss.NewStyle().Foreground(lipgloss.Color("#eba0ac")),
	mauve:     lipgloss.NewStyle().Foreground(lipgloss.Color("#cba6f7")),
	peach:     lipgloss.NewStyle().Foreground(lipgloss.Color("#fab387")),
	pink:      lipgloss.NewStyle().Foreground(lipgloss.Color("#f5c2e7")),
	red:       lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8")),
	lavender:  lipgloss.NewStyle().Foreground(lipgloss.Color("#b4befe")),
	rosewater: lipgloss.NewStyle().Foreground(lipgloss.Color("#f5e0dc")),
	sapphire:  lipgloss.NewStyle().Foreground(lipgloss.Color("#74c7ec")),
	sky:       lipgloss.NewStyle().Foreground(lipgloss.Color("#89dceb")),
	teal:      lipgloss.NewStyle().Foreground(lipgloss.Color("#94e2d5")),
	yellow:    lipgloss.NewStyle().Foreground(lipgloss.Color("#f9e2af")),
}

type paddingStylesStruct struct {
	smallAll lipgloss.Style
}

var paddingStyle = paddingStylesStruct{
	smallAll: lipgloss.NewStyle().Padding(1),
}
