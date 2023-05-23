package controller

import (
	"github.com/agusalex/golang-employee-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TagServiceMock struct {
	mock.Mock
}

func (m *TagServiceMock) GetAllTags() ([]models.Tag, error) {
	args := m.Called()
	return args.Get(0).([]models.Tag), args.Error(1)
}

func (m *TagServiceMock) GetTagByName(name string) (*models.Tag, error) {
	args := m.Called(name)
	return args.Get(0).(*models.Tag), args.Error(1)
}

func TestGetTagsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockTagService := new(TagServiceMock)
	mockTags := []models.Tag{{Name: "tag1"}, {Name: "tag2"}}
	mockTagService.On("GetAllTags").Return(mockTags, nil)

	TagService = mockTagService

	router := gin.Default()
	router.GET("/tags", GetTagsHandler)

	req, err := http.NewRequest(http.MethodGet, "/tags", nil)
	if err != nil {
		t.Fatalf("Failed to construct request: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "tag1")
	assert.Contains(t, rr.Body.String(), "tag2")

	mockTagService.AssertExpectations(t) // Checks if all mocked services were called as expected
}
