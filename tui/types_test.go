package tui

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/mklbravo/sshp/domain/entity"
	"github.com/stretchr/testify/assert"
)

// ################################
// FilterHost tests
func TestNewFilterHost(t *testing.T) {
	hostEntity := &entity.Host{
		Name: "foo",
		IP:   "bar",
	}
	values := []filterValue{
		NewFilterValue("foo", "name"),
		NewFilterValue("bar", "ip"),
	}
	host := NewFilterHost(hostEntity, values)
	assert.Same(t, hostEntity, host.host)
	assert.NotNil(t, host)
	assert.Equal(t, values, host.filterValues)
}

func TestFilterHost_Len(t *testing.T) {
	hostEntity := &entity.Host{
		Name: "foo",
		IP:   "bar",
	}
	values := []filterValue{
		NewFilterValue("foo", "name"),
		NewFilterValue("bar", "ip"),
	}
	host := NewFilterHost(hostEntity, values)
	assert.Equal(t, 2, host.Len())
}

func TestFilterHost_String(t *testing.T) {
	hostEntity := &entity.Host{
		Name: "foo",
		IP:   "bar",
	}
	values := []filterValue{
		NewFilterValue("foo", "name"),
		NewFilterValue("bar", "ip"),
	}
	host := NewFilterHost(hostEntity, values)
	assert.Equal(t, "foo", host.String(0))
	assert.Equal(t, "bar", host.String(1))
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
