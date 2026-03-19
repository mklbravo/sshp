package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/sahilm/fuzzy"
)

// ###################################
type filterList []*filterProfile

func NewFilterListFromHostEntities(profiles []*entity.Profile) filterList {
	filterProfiles := make(filterList, len(profiles))
	for index, entityProfile := range profiles {
		filterValues := []filterValue{
			NewFilterValue(string(entityProfile.Name), "name"),
			NewFilterValue(string(entityProfile.IP), "ip"),
			NewFilterValue(string(entityProfile.Username), "username"),
		}
		for _, profileDetails := range entityProfile.Details {
			filterValues = append(filterValues, NewFilterValue(profileDetails, "detail"))
		}

		filterProfiles[index] = &filterProfile{
			profile:      entityProfile,
			filterValues: filterValues,
		}
	}
	return filterProfiles
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

			filteredList = append(filteredList, NewFilterProfile(fh.profile, highlightedFilterValues))
		}
	}

	return filteredList
}

// ###################################
type filterProfile struct {
	profile      *entity.Profile
	filterValues []filterValue
}

func (this filterProfile) Len() int {
	return len(this.filterValues)
}

func (this filterProfile) String(index int) string {
	return this.filterValues[index].value
}

func NewFilterProfile(
	profile *entity.Profile,
	filterValues []filterValue,
) *filterProfile {
	return &filterProfile{
		profile:      profile,
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
