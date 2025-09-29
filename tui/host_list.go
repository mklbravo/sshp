package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/mklbravo/sshp/application"
	"github.com/mklbravo/sshp/domain/entity"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	hosts         []*entity.Host
	selectedIndex int
	textInput     textinput.Model
}

func NewHostListView(hostListUseCase *application.HostListUseCase) model {
	hosts, _ := hostListUseCase.Execute()
	// TODO: handle error

	// Initialize text input
	textInput := textinput.New()
	textInput.Focus()
	textInput.Placeholder = "Type to filter hosts..."
	textInput.Prompt = "  "
	textInput.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#74c7ec"))
	textInput.Width = 50

	return model{
		hosts:         hosts,
		selectedIndex: 0,
		textInput:     textInput,
	}
}

func (this model) Init() tea.Cmd {
	return textinput.Blink
}

func (this model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var inputCmd tea.Cmd

	this.textInput, inputCmd = this.textInput.Update(msg)

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keyMap.Down):
			this.SelectNext()
		case key.Matches(msg, keyMap.Up):
			this.SelectPrevious()
		case key.Matches(msg, keyMap.Quit):
			return this, tea.Quit
		}
	}

	return this, inputCmd
}

func (this model) View() string {
	if len(this.hosts) == 0 {
		return "No hosts available.\n"
	}

	selectedIndicatorStyle := lipgloss.NewStyle().
		Foreground(
			lipgloss.Color("#cba6f7"),
		)

	// Render the text input
	result := this.textInput.View() + "\n\n"

	// Render the list of hosts
	for index, host := range this.hosts {

		if index == this.selectedIndex {
			result += selectedIndicatorStyle.Render("󰁕 ")
		} else {
			result += "  "
		}

		result += fmt.Sprintf("%s (%s@%s:%d)\n", host.Name, host.Username, host.IP, host.Port)
	}

	result += lipgloss.NewStyle().Foreground(lipgloss.Color("#585b70")).Render("\nPress Esc or Ctrl+C to exit.\n")
	return result
}

func (this *model) SelectNext() {
	if len(this.hosts) > 0 {
		this.selectedIndex = (this.selectedIndex + 1) % len(this.hosts)
	}
}
func (this *model) SelectPrevious() {
	if len(this.hosts) > 0 {
		this.selectedIndex = (this.selectedIndex - 1 + len(this.hosts)) % len(this.hosts)
	}
}
