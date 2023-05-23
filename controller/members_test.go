package controller

import (
	"bytes"
	"encoding/json"
	"github.com/agusalex/golang-employee-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockMemberService struct {
	mock.Mock
}

func (m *MockMemberService) CreateMember(member models.Member) (models.Member, error) {
	args := m.Called(member)
	return args.Get(0).(models.Member), args.Error(1)
}

func (m *MockMemberService) UpdateMember(id string, member models.Member) (models.Member, error) {
	args := m.Called(id, member)
	return args.Get(0).(models.Member), args.Error(1)
}

func (m *MockMemberService) DeleteMember(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockMemberService) GetAllMembers() ([]models.Member, error) {
	args := m.Called()
	return args.Get(0).([]models.Member), args.Error(1)
}

func (m *MockMemberService) GetMemberByID(id string) (models.Member, error) {
	args := m.Called(id)
	return args.Get(0).(models.Member), args.Error(1)
}

func TestGetMemberHandler(t *testing.T) {
	memberService := new(MockMemberService)
	MemberService = memberService

	expectedMember := models.Member{
		Name: "John Doe",
		Type: "EMPLOYEE",
		Role: "Engineer",
	}

	memberService.On("GetMemberByID", "1").Return(expectedMember, nil)

	router := gin.Default()
	router.GET("/members/:id", GetMemberHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/members/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}

func TestGetMembersHandler(t *testing.T) {
	memberService := new(MockMemberService)
	MemberService = memberService

	expectedMembers := []models.Member{
		{
			Name: "John Doe",
			Type: "EMPLOYEE",
			Role: "Engineer",
		},
		{
			Name: "Jane Doe",
			Type: "EMPLOYEE",
			Role: "Engineer",
		},
	}

	memberService.On("GetAllMembers").Return(expectedMembers, nil)

	router := gin.Default()
	router.GET("/members", GetMembersHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/members", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
	assert.Contains(t, w.Body.String(), "Jane Doe")
}
func TestPostMembersHandler(t *testing.T) {
	memberService := new(MockMemberService)
	MemberService = memberService

	member := models.Member{
		Name:             "John Doe",
		Type:             "CONTRACTOR",
		ContractDuration: 1,
	}

	expectedMember := member
	expectedMember.ID = 1

	memberService.On("CreateMember", member).Return(expectedMember, nil)

	router := gin.Default()
	router.POST("/members", PostMembersHandler)

	jsonMember, _ := json.Marshal(member)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/members", bytes.NewBuffer(jsonMember))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}

func TestDeleteMemberHandler(t *testing.T) {
	memberService := new(MockMemberService)
	MemberService = memberService

	memberService.On("DeleteMember", "1").Return(nil)

	router := gin.Default()
	router.DELETE("/members/:id", DeleteMemberHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/members/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "member 1 deleted successfully")
}

func TestPutMemberHandler(t *testing.T) {
	memberService := new(MockMemberService)
	MemberService = memberService

	member := models.Member{
		Name:             "John Doe",
		Type:             "CONTRACTOR",
		ContractDuration: 1,
	}

	updatedMember := member
	updatedMember.Name = "Updated John Doe"

	memberService.On("UpdateMember", "1", member).Return(updatedMember, nil)

	router := gin.Default()
	router.PUT("/members/:id", PutMemberHandler)

	jsonMember, _ := json.Marshal(member)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/members/1", bytes.NewBuffer(jsonMember))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated John Doe")
}
