package tui

import (
	"testing"

	"github.com/mklbravo/sshp/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestModel_GetSelectedHost_NotSubmitted(t *testing.T) {
	model := Model{
		isSubmitted: false,
	}
	assert.Nil(t, model.GetSelectedHost())
}

func TestModel_GetSelectedHost_Submitted(t *testing.T) {
	profile := &entity.Profile{Name: "test.com"}
	filterProfile := NewFilterProfile(profile, []filterValue{
		NewFilterValue("test.com", "name"),
	})
	model := Model{
		matchedProfiles: filterList{filterProfile},
		selectedIndex:   0,
		isSubmitted:     true,
	}
	selected := model.GetSelectedHost()
	assert.Equal(t, profile, selected)
}

func TestModel_selectNext(t *testing.T) {
	profile1 := &entity.Profile{Name: "test1.com"}
	profile2 := &entity.Profile{Name: "test2.com"}
	filterProfile1 := NewFilterProfile(profile1, []filterValue{NewFilterValue("test1.com", "name")})
	filterProfile2 := NewFilterProfile(profile2, []filterValue{NewFilterValue("test2.com", "name")})
	model := Model{
		matchedProfiles: filterList{filterProfile1, filterProfile2},
		selectedIndex:   0,
	}

	model.selectNext()
	assert.Equal(t, 1, model.selectedIndex)

	model.selectNext()
	assert.Equal(t, 0, model.selectedIndex) // wrap around
}

func TestModel_selectPrevious(t *testing.T) {
	profile1 := &entity.Profile{Name: "test1.com"}
	profile2 := &entity.Profile{Name: "test2.com"}
	filterProfile1 := NewFilterProfile(profile1, []filterValue{NewFilterValue("test1.com", "name")})
	filterProfile2 := NewFilterProfile(profile2, []filterValue{NewFilterValue("test2.com", "name")})
	model := Model{
		matchedProfiles: filterList{filterProfile1, filterProfile2},
		selectedIndex:   0,
	}

	model.selectPrevious()
	assert.Equal(t, 1, model.selectedIndex)

	model.selectPrevious()
	assert.Equal(t, 0, model.selectedIndex) // wrap around
}

func TestModel_selectNext_EmptyList(t *testing.T) {
	model := Model{
		matchedProfiles: filterList{},
		selectedIndex:   0,
	}
	model.selectNext()
	assert.Equal(t, 0, model.selectedIndex) // no change
}

func TestModel_selectPrevious_EmptyList(t *testing.T) {
	model := Model{
		matchedProfiles: filterList{},
		selectedIndex:   0,
	}
	model.selectPrevious()
	assert.Equal(t, 0, model.selectedIndex) // no change
}
