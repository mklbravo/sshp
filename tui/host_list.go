package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/mklbravo/sshp/application"
	"github.com/mklbravo/sshp/domain/entity"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textInput textinput.Model
	hosts     []*entity.Host
	selected  *entity.Host
}

func NewHostListView(hostListUseCase *application.HostListUseCase) model {
	hosts, _ := hostListUseCase.Execute()
	// TODO: handle error
	// Initialize text input
	textInput := textinput.New()
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

	result := this.textInput.View() + "\n\n"
	result += "Available Hosts:\n"
	for _, host := range this.hosts {
		result += fmt.Sprintf("%d. %s (%s@%s:%d)\n", host.ID, host.Name, host.Username, host.IP, host.Port)
	}

	result += "\nPress Esc or Ctrl+C to exit.\n"
	return result
}
