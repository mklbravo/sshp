package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/mklbravo/sshp/application"
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/sahilm/fuzzy"

	tea "github.com/charmbracelet/bubbletea"
)

type filterableHostList []*entity.Host

func (this filterableHostList) String(index int) string {
	host := this[index]
	return fmt.Sprintf("%s %s %s", host.Name, host.Username, host.IP)
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
	return tea.Batch(tea.ClearScreen, textinput.Blink)
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
			return this, tea.Batch(tea.ClearScreen, tea.Quit)

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

			this.selectedIndex = 0
		}
	}

	return this, inputCmd
}

func (this Model) View() string {
	// Render the text input
	result := paddingStyle.smallAll.Render(this.textInput.View())

	if len(this.filteredHosts) == 0 {
		result += paddingStyle.smallAll.Render("No hosts found.\n")
	}

	hostTable := table.New().Border(lipgloss.HiddenBorder())

	for index, host := range this.filteredHosts {

		selectionPrefix := ""
		if index == this.selectedIndex {
			selectionPrefix = "󰁕 "
		}

		hostTable.Row(
			colorStyle.mauve.Render(selectionPrefix),
			colorStyle.sapphire.Render("󰍹 "),
			string(host.Name),
			colorStyle.teal.Render(" "),
			string(host.Username),
			colorStyle.sky.Render(" "),
			string(host.IP),
			colorStyle.sky.Render(" "),
			host.GetDetailsString(),
		)
	}

	result += hostTable.Render()

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
