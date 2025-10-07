package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/mklbravo/sshp/application"
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/sahilm/fuzzy"

	tea "github.com/charmbracelet/bubbletea"
)

type filterableHostList []*entity.Host

func (this filterableHostList) String(index int) string {
	return string(this[index].Name)
}
func (this filterableHostList) Len() int {
	return len(this)
}

func (this filterableHostList) GetFiltered(matches []fuzzy.Match) []*entity.Host {
	if len(matches) == 0 {
		return []*entity.Host{}
	}

	var filteredHosts []*entity.Host
	for _, match := range matches {
		filteredHosts = append(filteredHosts, this[match.Index])
	}
	return filteredHosts
}

type Model struct {
	filteredHosts      []*entity.Host
	filterableHostList filterableHostList
	isSubmitted        bool
	selectedIndex      int
	textInput          textinput.Model
}

func NewHostListView(hostListUseCase *application.HostListUseCase) Model {
	hosts, _ := hostListUseCase.Execute()
	// TODO: handle error

	// Initialize text input
	textInput := textinput.New()
	textInput.Focus()
	textInput.Placeholder = "Type to filter hosts..."
	textInput.Prompt = "  "
	textInput.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#74c7ec"))
	textInput.Width = 50

	return Model{
		filterableHostList: hosts,
		filteredHosts:      hosts,
		isSubmitted:        false,
		selectedIndex:      0,
		textInput:          textInput,
	}
}

func (this Model) Init() tea.Cmd {
	return textinput.Blink
}

func (this Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var inputCmd tea.Cmd

	this.textInput, inputCmd = this.textInput.Update(msg)

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keyMap.Down):
			this.selectNext()

		case key.Matches(msg, keyMap.Submit):
			this.isSubmitted = true
			return this, tea.Quit

		case key.Matches(msg, keyMap.Quit):
			return this, tea.Quit

		case key.Matches(msg, keyMap.Up):
			this.selectPrevious()
		default:
			if this.textInput.Value() == "" {
				this.filteredHosts = this.filterableHostList
			} else {
				matches := fuzzy.FindFrom(this.textInput.Value(), this.filterableHostList)
				this.filteredHosts = this.filterableHostList.GetFiltered(matches)
			}
		}
	}

	return this, inputCmd
}

func (this Model) View() string {
	// Render the text input
	result := this.textInput.View() + "\n\n"

	if len(this.filteredHosts) == 0 {
		result += "No hosts...\n"
	}

	// Render the list of hosts
	for index, host := range this.filteredHosts {

		if index == this.selectedIndex {
			result += styles.mauve.Render("󰁕 ")
		} else {
			result += "  "
		}

		result += fmt.Sprintf("%s%s\t%s%s\t%s%s\n",
			styles.sapphire.Render("󰍹  "),
			host.Name,
			styles.teal.Render(" "),
			host.Username,
			styles.sky.Render(" "),
			host.IP,
		)
		// result += fmt.Sprintf("󰍹  %s  %s  %s\n", host.Name, host.Username, host.IP)
	}

	result += lipgloss.NewStyle().Foreground(lipgloss.Color("#585b70")).Render("\nPress Esc or Ctrl+C to exit.\n")
	return result
}

func (this *Model) GetSelectedHost() *entity.Host {
	if !this.isSubmitted {
		return nil
	}

	return this.filteredHosts[this.selectedIndex]
}

func (this *Model) selectNext() {
	if len(this.filteredHosts) > 0 {
		this.selectedIndex = (this.selectedIndex + 1) % len(this.filteredHosts)
	}
}
func (this *Model) selectPrevious() {
	if len(this.filteredHosts) > 0 {
		this.selectedIndex = (this.selectedIndex - 1 + len(this.filteredHosts)) % len(this.filteredHosts)
	}
}
