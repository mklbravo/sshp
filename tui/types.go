package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// ###################################
type filterValue struct {
	value            string
	column           string
	highlightIndexes []int
}

func NewFilterValue(value string, column string) filterValue {
	return filterValue{
		value:            value,
		column:           column,
		highlightIndexes: []int{},
	}
}

func (this filterValue) AddHighlightIndex(index ...int) filterValue {
	return filterValue{
		value:            this.value,
		column:           this.column,
		highlightIndexes: index,
	}
}

func (this filterValue) GetHightlightedString(hightlightStyle lipgloss.Style) string {
	if len(this.highlightIndexes) == 0 {
		return this.value
	}

	var ranges []lipgloss.Range
	for _, index := range this.highlightIndexes {
		ranges = append(ranges, lipgloss.NewRange(index, index+1, hightlightStyle))
	}

	return lipgloss.StyleRanges(
		this.value,
		ranges...,
	)
}
