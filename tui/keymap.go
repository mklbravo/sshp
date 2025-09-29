package tui

import "github.com/charmbracelet/bubbles/key"

type keyMapStruct struct {
	Down key.Binding
	Quit key.Binding
	Up   key.Binding
}

var keyMap = keyMapStruct{
	Down: key.NewBinding(
		key.WithKeys("down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
	),
}
