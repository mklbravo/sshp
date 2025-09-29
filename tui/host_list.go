package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/mklbravo/sshp/application"
	"github.com/mklbravo/sshp/domain/entity"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	hosts     []*entity.Host
	selected  *entity.Host
	textInput textinput.Model
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
		hosts:     hosts,
		selected:  nil,
		textInput: textInput,
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
		switch msg.String() {
		case "esc", "ctrl+c":
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
	for i, host := range this.hosts {

		if i == 1 {
			result += selectedIndicatorStyle.Render("󰁕 ")
		} else {
			result += "  "
		}

		result += fmt.Sprintf("%s (%s@%s:%d)\n", host.Name, host.Username, host.IP, host.Port)
	}

	result += lipgloss.NewStyle().Foreground(lipgloss.Color("#585b70")).Render("\nPress Esc or Ctrl+C to exit.\n")
	return result
}
