package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/agusalex/golang-employee-crud/models"
)

type MockTagRepository struct {
	mock.Mock
}

func (m *MockTagRepository) GetAll() ([]models.Tag, error) {
	args := m.Called()
	return args.Get(0).([]models.Tag), args.Error(1)
}

func (m *MockTagRepository) GetByName(name string) (*models.Tag, error) {
	args := m.Called(name)
	return args.Get(0).(*models.Tag), args.Error(1)
}
func TestTagService_GetAllTags(t *testing.T) {
	mockRepo := new(MockTagRepository)
	tagService := NewTagService(mockRepo)

	tags := []models.Tag{{Name: "Tag1"}, {Name: "Tag2"}}

	mockRepo.On("GetAll").Return(tags, nil)

	result, err := tagService.GetAllTags()

	assert.NoError(t, err)
	assert.Equal(t, tags, result)

	mockRepo.AssertExpectations(t)
}

func TestTagService_GetTagByName(t *testing.T) {
	mockRepo := new(MockTagRepository)
	tagService := NewTagService(mockRepo)

	tag := &models.Tag{Name: "Tag1"}

	mockRepo.On("GetByName", "Tag1").Return(tag, nil)

	result, err := tagService.GetTagByName("Tag1")

	assert.NoError(t, err)
	assert.Equal(t, tag, result)

	mockRepo.AssertExpectations(t)
}
