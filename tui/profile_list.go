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

var highlightStyle = colorStyle.mauve.Underline(true).Bold(true)

type Model struct {
	filterProfiles  filterList
	matchedProfiles filterList
	isSubmitted     bool
	selectedIndex   int
	textInput       textinput.Model
}

func NewProfileListView(profileListUseCase *application.ProfileListUseCase) Model {
	// Initialize text input
	textInput := textinput.New()
	textInput.Focus()
	textInput.Placeholder = "Type to filter profiles..."
	textInput.Prompt = "  "
	textInput.PromptStyle = colorStyle.sapphire
	textInput.Width = 50

	profiles, _ := profileListUseCase.Execute()
	// TODO: handle error

	filterProfiles := NewFilterListFromHostEntities(profiles)

	return Model{
		filterProfiles:  filterProfiles,
		isSubmitted:     false,
		matchedProfiles: filterProfiles, // Initially, all profiles are matched
		selectedIndex:   0,
		textInput:       textInput,
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
				this.matchedProfiles = this.filterProfiles
			} else {
				this.matchedProfiles = this.filterProfiles.Filter(this.textInput.Value())
			}

			this.selectedIndex = 0
		}
	}

	return this, inputCmd
}

func (this Model) View() string {
	// Render the text input
	result := paddingStyle.smallAll.Render(this.textInput.View())

	if len(this.matchedProfiles) == 0 {
		result += paddingStyle.smallAll.Render("No profiles found.\n")
	}

	profileTable := table.New().Border(lipgloss.HiddenBorder())

	for index, filterHost := range this.matchedProfiles {
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

		profileTable.Row(
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

	result += profileTable.Render()

	result += lipgloss.NewStyle().Foreground(lipgloss.Color("#585b70")).Render("\nPress Esc or Ctrl+C to exit.\n")
	return result
}

func (this *Model) GetSelectedHost() *entity.Profile {
	if !this.isSubmitted {
		return nil
	}

	return this.matchedProfiles[this.selectedIndex].host
}

func (this *Model) selectNext() {
	if len(this.matchedProfiles) > 0 {
		this.selectedIndex = (this.selectedIndex + 1) % len(this.matchedProfiles)
	}
}

func (this *Model) selectPrevious() {
	if len(this.matchedProfiles) > 0 {
		this.selectedIndex = (this.selectedIndex - 1 + len(this.matchedProfiles)) % len(this.matchedProfiles)
	}
}
