package tui

import "github.com/charmbracelet/bubbles/key"

type keyMapStruct struct {
	Quit key.Binding
}

var keyMap = keyMapStruct{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
	),
}
