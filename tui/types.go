package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/sahilm/fuzzy"
)

// ###################################
type filterList []*filterHost

func NewFilterListFromHostEntities(hosts []*entity.Host) filterList {
	filterHosts := make(filterList, len(hosts))
	for index, entityHost := range hosts {
		filterValues := []filterValue{
			NewFilterValue(string(entityHost.Name), "name"),
			NewFilterValue(string(entityHost.IP), "ip"),
			NewFilterValue(string(entityHost.Username), "username"),
		}
		for _, hostDetails := range entityHost.Details {
			filterValues = append(filterValues, NewFilterValue(hostDetails, "detail"))
		}

		filterHosts[index] = &filterHost{
			host:         entityHost,
			filterValues: filterValues,
		}
	}
	return filterHosts
}

// Returns a new filterList based on current containing only the filterHosts that match the query
func (this *filterList) Filter(query string) filterList {

	filteredList := make(filterList, 0)
	for _, fh := range *this {
		match := fuzzy.FindFrom(query, fh)
		if len(match) > 0 {
			matchMap := make(map[int]fuzzy.Match)
			for _, m := range match {
				matchMap[m.Index] = m
			}

			highlightedFilterValues := make([]filterValue, len(fh.filterValues))
			for index, filterValue := range fh.filterValues {
				if _, ok := matchMap[index]; ok {
					highlightedFilterValues[index] = filterValue.AddHighlightIndex(matchMap[index].MatchedIndexes...)
				} else {
					highlightedFilterValues[index] = filterValue
				}
			}

			filteredList = append(filteredList, NewFilterHost(fh.host, highlightedFilterValues))
		}
	}

	return filteredList
}

// ###################################
type filterHost struct {
	host         *entity.Host
	filterValues []filterValue
}

func (this filterHost) Len() int {
	return len(this.filterValues)
}

func (this filterHost) String(index int) string {
	return this.filterValues[index].value
}

func NewFilterHost(
	host *entity.Host,
	filterValues []filterValue,
) *filterHost {
	return &filterHost{
		host:         host,
		filterValues: filterValues,
	}
}

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
