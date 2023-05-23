package services

import (
	"github.com/agusalex/golang-employee-crud/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockSearchRepository struct {
	mock.Mock
}

func (m *MockSearchRepository) SearchMembers(tags []string, memberType string) ([]models.Member, error) {
	args := m.Called(tags, memberType)
	return args.Get(0).([]models.Member), args.Error(1)
}

func TestSearchService_SearchMembers(t *testing.T) {
	mockRepo := new(MockSearchRepository)
	searchService := NewSearchService(mockRepo)

	members := []models.Member{
		{Name: "Member1", Type: "Type1"},
		{Name: "Member2", Type: "Type2"},
	}

	mockRepo.On("SearchMembers", []string{"tag1", "tag2"}, "Type1").Return(members, nil)

	result, err := searchService.SearchMembers([]string{"tag1", "tag2"}, "Type1")

	assert.NoError(t, err)
	assert.Equal(t, members, result)

	mockRepo.AssertExpectations(t)
}
