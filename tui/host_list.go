package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/mklbravo/sshp/application"
	"github.com/mklbravo/sshp/domain/entity"

	tea "github.com/charmbracelet/bubbletea"
)

var highlightStyle = colorStyle.yellow.Bold(true)

type Model struct {
	filterHosts   filterList
	matchedHosts  filterList
	isSubmitted   bool
	selectedIndex int
	textInput     textinput.Model
}

func NewHostListView(hostListUseCase *application.HostListUseCase) Model {
	// Initialize text input
	textInput := textinput.New()
	textInput.Focus()
	textInput.Placeholder = "Type to filter hosts..."
	textInput.Prompt = "  "
	textInput.PromptStyle = colorStyle.sapphire
	textInput.Width = 50

	hosts, _ := hostListUseCase.Execute()
	// TODO: handle error

	filterHosts := NewFilterListFromHostEntities(hosts)

	return Model{
		filterHosts:   filterHosts,
		isSubmitted:   false,
		matchedHosts:  filterHosts, // Initially, all hosts are matched
		selectedIndex: 0,
		textInput:     textInput,
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
				this.matchedHosts = this.filterHosts
			} else {
				this.matchedHosts = this.filterHosts.Filter(this.textInput.Value())
			}

			this.selectedIndex = 0
		}
	}

	return this, inputCmd
}

func (this Model) View() string {
	// Render the text input
	result := paddingStyle.smallAll.Render(this.textInput.View())

	if len(this.matchedHosts) == 0 {
		result += paddingStyle.smallAll.Render("No hosts found.\n")
	}

	hostTable := table.New().Border(lipgloss.HiddenBorder())

	for index, filterHost := range this.matchedHosts {
		selectionColumnContent := ""
		if index == this.selectedIndex {
			selectionColumnContent = colorStyle.mauve.Render("󰁕 ")
		}

		nameColumnStringBuilder := &strings.Builder{}
		nameColumnStringBuilder.WriteString(colorStyle.sapphire.Render("󰍹 "))

		usernameColumnStringBuilder := &strings.Builder{}
		usernameColumnStringBuilder.WriteString(colorStyle.teal.Render(" "))

		ipColumnStringBuilder := &strings.Builder{}
		ipColumnStringBuilder.WriteString(colorStyle.sky.Render(" "))

		detailColumnStringBuilder := &strings.Builder{}

		for _, filterValue := range filterHost.filterValues {
			switch filterValue.column {
			case "name":
				nameColumnStringBuilder.WriteString(filterValue.GetHightlightedString(highlightStyle))
			case "username":
				usernameColumnStringBuilder.WriteString(filterValue.GetHightlightedString(highlightStyle))
			case "ip":
				ipColumnStringBuilder.WriteString(filterValue.GetHightlightedString(highlightStyle))
			case "detail":
				if detailColumnStringBuilder.Len() > 0 {
					detailColumnStringBuilder.WriteString(" | ")
				}
				detailColumnStringBuilder.WriteString(filterValue.GetHightlightedString(highlightStyle))
			}
		}

		// Do not show the icon if there is no details
		detailsColumnContent := ""
		if detailColumnStringBuilder.Len() > 0 {
			detailsColumnContent = fmt.Sprintf("%s %s",
				colorStyle.sky.Render(" "),
				detailColumnStringBuilder.String(),
			)
		}

		hostTable.Row(
			selectionColumnContent,
			nameColumnStringBuilder.String(),
			" ", // Spacer column
			usernameColumnStringBuilder.String(),
			" ", // Spacer column
			ipColumnStringBuilder.String(),
			" ", // Spacer column
			detailsColumnContent,
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

	return this.matchedHosts[this.selectedIndex].host
}

func (this *Model) selectNext() {
	if len(this.matchedHosts) > 0 {
		this.selectedIndex = (this.selectedIndex + 1) % len(this.matchedHosts)
	}
}

func (this *Model) selectPrevious() {
	if len(this.matchedHosts) > 0 {
		this.selectedIndex = (this.selectedIndex - 1 + len(this.matchedHosts)) % len(this.matchedHosts)
	}
}
