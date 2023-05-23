package services

import (
	"github.com/agusalex/golang-employee-crud/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockMemberRepository struct {
	mock.Mock
}

func (m *MockMemberRepository) GetByID(id string) (models.Member, error) {
	args := m.Called(id)
	return args.Get(0).(models.Member), args.Error(1)
}

func (m *MockMemberRepository) Save(member models.Member) (models.Member, error) {
	args := m.Called(member)
	return args.Get(0).(models.Member), args.Error(1)
}

func (m *MockMemberRepository) Update(existingMember models.Member, member models.Member) (models.Member, error) {
	args := m.Called(existingMember, member)
	return args.Get(0).(models.Member), args.Error(1)
}

func (m *MockMemberRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockMemberRepository) GetAll() ([]models.Member, error) {
	args := m.Called()
	return args.Get(0).([]models.Member), args.Error(1)
}

func TestMemberService_GetAllMembers(t *testing.T) {
	mockMemberRepo := new(MockMemberRepository)
	mockTagRepo := new(MockTagRepository)
	memberService := NewMemberService(mockMemberRepo, mockTagRepo)

	members := []models.Member{
		{
			Name:             "John Doe",
			Type:             "CONTRACTOR",
			Tags:             nil,
			ContractDuration: 1,
		},
		{
			Name:             "John Doe",
			Type:             "CONTRACTOR",
			Tags:             nil,
			ContractDuration: 1,
		},
	}

	mockMemberRepo.On("GetAll").Return(members, nil)

	result, err := memberService.GetAllMembers()

	assert.NoError(t, err)
	assert.Equal(t, members, result)

	mockMemberRepo.AssertExpectations(t)
}

func TestMemberService_GetMemberByID(t *testing.T) {
	mockMemberRepo := new(MockMemberRepository)
	mockTagRepo := new(MockTagRepository)
	memberService := NewMemberService(mockMemberRepo, mockTagRepo)

	member := models.Member{
		Name:             "John Doe",
		Type:             "CONTRACTOR",
		Tags:             nil,
		ContractDuration: 1,
	}

	mockMemberRepo.On("GetByID", "1").Return(member, nil)

	result, err := memberService.GetMemberByID("1")

	assert.NoError(t, err)
	assert.Equal(t, member, result)

	mockMemberRepo.AssertExpectations(t)
}
func TestMemberService_CreateMember(t *testing.T) {
	mockMemberRepo := new(MockMemberRepository)
	mockTagRepo := new(MockTagRepository)
	memberService := NewMemberService(mockMemberRepo, mockTagRepo)

	tag := models.Tag{Name: "Tag1"}
	member := models.Member{
		Name:             "John Doe",
		Type:             "CONTRACTOR",
		ContractDuration: 1,
		Tags:             []models.Tag{tag},
	}

	mockTagRepo.On("GetByName", "Tag1").Return(&tag, nil)
	mockMemberRepo.On("Save", member).Return(member, nil)

	result, err := memberService.CreateMember(member)

	assert.NoError(t, err)
	assert.Equal(t, member, result)

	mockMemberRepo.AssertExpectations(t)
	mockTagRepo.AssertExpectations(t)
}

func TestMemberService_UpdateMember(t *testing.T) {
	mockMemberRepo := new(MockMemberRepository)
	mockTagRepo := new(MockTagRepository)
	memberService := NewMemberService(mockMemberRepo, mockTagRepo)

	tag := models.Tag{Name: "Tag1"}
	member := models.Member{
		Name:             "John Doe",
		Type:             "CONTRACTOR",
		ContractDuration: 1,
		Tags:             []models.Tag{tag},
	}

	mockTagRepo.On("GetByName", "Tag1").Return(&tag, nil)
	mockMemberRepo.On("GetByID", "1").Return(member, nil)
	mockMemberRepo.On("Update", member, member).Return(member, nil)

	result, err := memberService.UpdateMember("1", member)

	assert.NoError(t, err)
	assert.Equal(t, member, result)

	mockMemberRepo.AssertExpectations(t)
	mockTagRepo.AssertExpectations(t)
}

func TestMemberService_DeleteMember(t *testing.T) {
	mockMemberRepo := new(MockMemberRepository)
	mockTagRepo := new(MockTagRepository)
	memberService := NewMemberService(mockMemberRepo, mockTagRepo)

	mockMemberRepo.On("Delete", "1").Return(nil)

	err := memberService.DeleteMember("1")

	assert.NoError(t, err)

	mockMemberRepo.AssertExpectations(t)
}
