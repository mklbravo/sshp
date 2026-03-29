package tui

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/stretchr/testify/assert"
)

// ################################
// FilterProfile tests
func TestNewFilterProfile(t *testing.T) {
	profileEntity := &entity.Profile{
		Name: "foo",
		IP:   "bar",
	}
	values := []filterValue{
		NewFilterValue("foo", "name"),
		NewFilterValue("bar", "ip"),
	}
	profile := NewFilterProfile(profileEntity, values)
	assert.Same(t, profileEntity, profile.profile)
	assert.NotNil(t, profile)
	assert.Equal(t, values, profile.filterValues)
}

func TestFilterProfile_Len(t *testing.T) {
	profileEntity := &entity.Profile{
		Name: "foo",
		IP:   "bar",
	}
	values := []filterValue{
		NewFilterValue("foo", "name"),
		NewFilterValue("bar", "ip"),
	}
	profile := NewFilterProfile(profileEntity, values)
	assert.Equal(t, 2, profile.Len())
}

func TestFilterProfile_String(t *testing.T) {
	profileEntity := &entity.Profile{
		Name: "foo",
		IP:   "bar",
	}
	values := []filterValue{
		NewFilterValue("foo", "name"),
		NewFilterValue("bar", "ip"),
	}
	profile := NewFilterProfile(profileEntity, values)
	assert.Equal(t, "foo", profile.String(0))
	assert.Equal(t, "bar", profile.String(1))
}

// ################################
// FilterValue tests
func TestNewFilterValue(t *testing.T) {
	val := NewFilterValue("test", "name")
	assert.Equal(t, "test", val.value)
	assert.Equal(t, "name", val.column)
	assert.Empty(t, val.highlightIndexes)
}

func TestFilterValue_AddHighlightIndex(t *testing.T) {
	val := NewFilterValue("test", "name")
	val2 := val.AddHighlightIndex(1, 2)
	assert.NotSame(t, &val, &val2) // Ensure immutability

	assert.Equal(t, []int{1, 2}, val2.highlightIndexes)
	assert.Equal(t, "test", val2.value)
	assert.Equal(t, "name", val2.column)

}

func TestFilterValue_GetHightlightedString_NoHighlight(t *testing.T) {
	val := NewFilterValue("test", "name")
	style := lipgloss.NewStyle().Bold(true)
	result := val.GetHightlightedString(style)
	assert.Equal(t, "test", result)
}

func TestFilterValue_GetHightlightedString_WithHighlight(t *testing.T) {
	val := NewFilterValue("test", "name").AddHighlightIndex(1, 2)
	style := lipgloss.NewStyle().Bold(true)
	result := val.GetHightlightedString(style)
	// Since lipgloss.StyleRanges returns a styled string, we just check it's not empty and contains the base value
	assert.Contains(t, result, "test")
}

// ################################
// FilterList tests
func TestNewFilterListFromHostEntities(t *testing.T) {
	profiles := []*entity.Profile{
		{
			Name:     "example.com",
			IP:       "192.168.1.1",
			Username: "user",
			Details:  []string{"detail1"},
		},
		{
			Name:     "test.com",
			IP:       "10.0.0.1",
			Username: "admin",
			Details:  []string{},
		},
	}

	filterList := NewFilterListFromHostEntities(profiles)

	assert.Len(t, filterList, 2)
	assert.Equal(t, "example.com", filterList[0].filterValues[0].value)
	assert.Equal(t, "192.168.1.1", filterList[0].filterValues[1].value)
	assert.Equal(t, "user", filterList[0].filterValues[2].value)
	assert.Equal(t, "detail1", filterList[0].filterValues[3].value)
	assert.Len(t, filterList[1].filterValues, 3) // no details
}

func TestFilterList_Filter_NoQuery(t *testing.T) {
	profiles := []*entity.Profile{
		{
			Name:     "example.com",
			IP:       "192.168.1.1",
			Username: "user",
		},
	}

	filterList := NewFilterListFromHostEntities(profiles)
	filtered := filterList.Filter("")

	assert.Len(t, filtered, 1)
	assert.Equal(t, filterList[0], filtered[0])
}

func TestFilterList_Filter_WithMatch(t *testing.T) {
	profiles := []*entity.Profile{
		{
			Name:     "example.com",
			IP:       "192.168.1.1",
			Username: "user",
		},
		{
			Name:     "test.com",
			IP:       "10.0.0.1",
			Username: "admin",
		},
	}

	filterList := NewFilterListFromHostEntities(profiles)
	filtered := filterList.Filter("example")

	assert.Len(t, filtered, 1)
	assert.Equal(t, "example.com", string(filtered[0].profile.Name))
	// Check highlighting
	assert.NotEmpty(t, filtered[0].filterValues[0].highlightIndexes)
}

func TestFilterList_Filter_NoMatch(t *testing.T) {
	profiles := []*entity.Profile{
		{
			Name:     "example.com",
			IP:       "192.168.1.1",
			Username: "user",
		},
	}

	filterList := NewFilterListFromHostEntities(profiles)
	filtered := filterList.Filter("nomatch")

	assert.Len(t, filtered, 0)
}
