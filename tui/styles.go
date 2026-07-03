package tui

import "github.com/charmbracelet/lipgloss"

// terminalColorStyles uses ANSI terminal colors (0-15) which adapt to the user's terminal theme.
type terminalColorStyles struct {
	highlight lipgloss.Style // Selected item, emphasis (typically bright magenta/purple)
	accent    lipgloss.Style // Input prompts, primary UI elements (typically blue)
	secondary lipgloss.Style // Secondary icons, less important elements (typically cyan)
	tertiary  lipgloss.Style // Tertiary elements (typically yellow)
	muted     lipgloss.Style // Deemphasized text, background elements (typically gray)
	error     lipgloss.Style // Error states (typically red)
	success   lipgloss.Style // Success/confirmation states (typically green)
}

var colorStyle = terminalColorStyles{
	highlight: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
	accent:    lipgloss.NewStyle().Foreground(lipgloss.Color("4")),
	secondary: lipgloss.NewStyle().Foreground(lipgloss.Color("6")),
	tertiary:  lipgloss.NewStyle().Foreground(lipgloss.Color("3")),
	muted:     lipgloss.NewStyle().Foreground(lipgloss.Color("8")),
	error:     lipgloss.NewStyle().Foreground(lipgloss.Color("1")),
	success:   lipgloss.NewStyle().Foreground(lipgloss.Color("2")),
}

type paddingStylesStruct struct {
	smallAll lipgloss.Style
}

var paddingStyle = paddingStylesStruct{
	smallAll: lipgloss.NewStyle().Padding(1),
}
