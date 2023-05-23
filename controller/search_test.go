package controller

import (
	"github.com/agusalex/golang-employee-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SearchServiceMock struct {
	mock.Mock
}

func (m *SearchServiceMock) SearchMembers(tags []string, memberType string) ([]models.Member, error) {
	args := m.Called(tags, memberType)
	return args.Get(0).([]models.Member), args.Error(1)
}

// FIXME These tests need improvement
func TestSearchMembersHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSearchService := new(SearchServiceMock)
	mockMembers := []models.Member{{Name: "John Doe", Type: "EMPLOYEE", Role: "Engineer"}, {Name: "Jane Doe", Type: "EMPLOYEE", Role: "Engineer"}}
	mockSearchService.On("SearchMembers", []string{"Java"}, "CONTRACTOR").Return(mockMembers, nil)

	SearchService = mockSearchService

	router := gin.Default()
	router.GET("/search/members", SearchMembersHandler)

	req, err := http.NewRequest(http.MethodGet, "/search/members?type=CONTRACTOR&tags=Java", nil)
	if err != nil {
		t.Fatalf("Failed to construct request: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "John Doe")
	assert.Contains(t, rr.Body.String(), "Jane Doe")

	mockSearchService.AssertExpectations(t)
}
