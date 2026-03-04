package tui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/mklbravo/sshp/application"
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/sahilm/fuzzy"

	tea "github.com/charmbracelet/bubbletea"
)

var highlightStyle = colorStyle.moccasin.Bold(true)

type hostMatch struct {
	host            *entity.Host
	score           int
	nameIndexes     []int
	usernameIndexes []int
	ipIndexes       []int
	detailIndexes   [][]int // one entry per detail item
}

type Model struct {
	allHosts      []*entity.Host
	matches       []hostMatch
	isSubmitted   bool
	selectedIndex int
	textInput     textinput.Model
}

func NewHostListView(hostListUseCase *application.HostListUseCase) Model {
	hosts, _ := hostListUseCase.Execute()
	// TODO: handle error

	// Initialize text input
	textInput := textinput.New()
	textInput.Focus()
	textInput.Placeholder = "Type to filter hosts..."
	textInput.Prompt = "  "
	textInput.PromptStyle = colorStyle.sapphire
	textInput.Width = 50

	return Model{
		allHosts:      hosts,
		matches:       allHostsAsMatches(hosts),
		isSubmitted:   false,
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
				this.matches = allHostsAsMatches(this.allHosts)
			} else {
				this.matches = searchHosts(this.textInput.Value(), this.allHosts)
			}

			this.selectedIndex = 0
		}
	}

	return this, inputCmd
}

func (this Model) View() string {
	// Render the text input
	result := paddingStyle.smallAll.Render(this.textInput.View())

	if len(this.matches) == 0 {
		result += paddingStyle.smallAll.Render("No hosts found.\n")
	}

	hostTable := table.New().Border(lipgloss.HiddenBorder())

	for index, m := range this.matches {
		host := m.host

		selectionColumnContent := ""
		if index == this.selectedIndex {
			selectionColumnContent = colorStyle.mauve.Render("󰁕 ")
		}

		nameColumnContent := fmt.Sprintf(
			"%s  %s", // Two spaces for padding
			colorStyle.sapphire.Render("󰍹"),
			renderWithHighlights(string(host.Name), m.nameIndexes),
		)

		usernameColumnContent := fmt.Sprintf(
			"%s %s",
			colorStyle.teal.Render(""),
			renderWithHighlights(string(host.Username), m.usernameIndexes),
		)

		ipColumnContent := fmt.Sprintf(
			"%s %s",
			colorStyle.sky.Render(""),
			renderWithHighlights(string(host.IP), m.ipIndexes),
		)

		detailsColumnContent := ""
		if host.HasDetails() {
			detailsColumnContent = fmt.Sprintf(
				"%s %s",
				colorStyle.sky.Render(""),
				buildDetailsString(host.Details, m.detailIndexes),
			)
		}

		hostTable.Row(
			selectionColumnContent,
			nameColumnContent,
			" ", // Spacer column
			usernameColumnContent,
			" ", // Spacer column
			ipColumnContent,
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

	return this.matches[this.selectedIndex].host
}

func (this *Model) selectNext() {
	if len(this.matches) > 0 {
		this.selectedIndex = (this.selectedIndex + 1) % len(this.matches)
	}
}

func (this *Model) selectPrevious() {
	if len(this.matches) > 0 {
		this.selectedIndex = (this.selectedIndex - 1 + len(this.matches)) % len(this.matches)
	}
}

func allHostsAsMatches(hosts []*entity.Host) []hostMatch {
	matches := make([]hostMatch, len(hosts))
	for i, h := range hosts {
		matches[i] = hostMatch{
			host:          h,
			detailIndexes: make([][]int, len(h.Details)),
		}
	}
	return matches
}

func searchHosts(query string, hosts []*entity.Host) []hostMatch {
	var results []hostMatch

	for _, host := range hosts {
		m := matchHost(query, host)
		if m.score > 0 {
			results = append(results, m)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].score > results[j].score
	})

	return results
}

func matchHost(query string, host *entity.Host) hostMatch {
	m := hostMatch{
		host:          host,
		detailIndexes: make([][]int, len(host.Details)),
	}

	if fm := firstMatch(query, string(host.Name)); fm != nil {
		m.nameIndexes = fm.MatchedIndexes
		if fm.Score > m.score {
			m.score = fm.Score
		}
	}

	if fm := firstMatch(query, string(host.Username)); fm != nil {
		m.usernameIndexes = fm.MatchedIndexes
		if fm.Score > m.score {
			m.score = fm.Score
		}
	}

	if fm := firstMatch(query, string(host.IP)); fm != nil {
		m.ipIndexes = fm.MatchedIndexes
		if fm.Score > m.score {
			m.score = fm.Score
		}
	}

	for i, detail := range host.Details {
		if fm := firstMatch(query, detail); fm != nil {
			m.detailIndexes[i] = fm.MatchedIndexes
			if fm.Score > m.score {
				m.score = fm.Score
			}
		}
	}

	return m
}

func firstMatch(query, s string) *fuzzy.Match {
	results := fuzzy.Find(query, []string{s})
	if len(results) == 0 {
		return nil
	}
	return &results[0]
}

func renderWithHighlights(s string, matchedIndexes []int) string {
	if len(matchedIndexes) == 0 {
		return s
	}

	indexSet := make(map[int]bool, len(matchedIndexes))
	for _, i := range matchedIndexes {
		indexSet[i] = true
	}

	var sb strings.Builder
	for i, ch := range s {
		if indexSet[i] {
			sb.WriteString(highlightStyle.Render(string(ch)))
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func buildDetailsString(details []string, detailIndexes [][]int) string {
	var sb strings.Builder

	for i, detail := range details {
		var indexes []int
		if i < len(detailIndexes) {
			indexes = detailIndexes[i]
		}

		sb.WriteString(renderWithHighlights(detail, indexes))

		if i != len(details)-1 {
			sb.WriteString(" | ")
		}
	}

	return sb.String()
}
