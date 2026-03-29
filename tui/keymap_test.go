package tui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyMap_Down(t *testing.T) {
	assert.Contains(t, keyMap.Down.Keys(), "down")
}

func TestKeyMap_Up(t *testing.T) {
	assert.Contains(t, keyMap.Up.Keys(), "up")
}

func TestKeyMap_Submit(t *testing.T) {
	assert.Contains(t, keyMap.Submit.Keys(), "enter")
}

func TestKeyMap_Quit(t *testing.T) {
	keys := keyMap.Quit.Keys()
	assert.Contains(t, keys, "ctrl+c")
	assert.Contains(t, keys, "esc")
}
