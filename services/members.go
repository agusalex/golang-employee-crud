package services

import (
	"github.com/agusalex/golang-employee-crud/models"
	repositories2 "github.com/agusalex/golang-employee-crud/repositories"
)

type MemberServiceInterface interface {
	GetAllMembers() ([]models.Member, error)
	GetMemberByID(id string) (models.Member, error)
	CreateMember(member models.Member) (models.Member, error)
	UpdateMember(id string, member models.Member) (models.Member, error)
	DeleteMember(id string) error
}

type MemberServiceStruct struct {
	memberRepo repositories2.MemberRepositoryInterface
	tagRepo    repositories2.TagRepositoryInterface
}

var MemberService MemberServiceInterface

func NewMemberService(memberRepo repositories2.MemberRepositoryInterface,
	tagRepo repositories2.TagRepositoryInterface) MemberServiceInterface {
	return &MemberServiceStruct{
		memberRepo: memberRepo,
		tagRepo:    tagRepo,
	}
}

func (s *MemberServiceStruct) GetAllMembers() ([]models.Member, error) {
	return s.memberRepo.GetAll()
}

func (s *MemberServiceStruct) GetMemberByID(id string) (models.Member, error) {
	return s.memberRepo.GetByID(id)
}

func (s *MemberServiceStruct) CreateMember(member models.Member) (models.Member, error) {
	for i, tag := range member.Tags {
		existingTag, err := s.tagRepo.GetByName(tag.Name)
		if err != nil {
			return models.Member{}, err
		}
		if existingTag != nil {
			member.Tags[i] = *existingTag
		}
	}
	return s.memberRepo.Save(member)
}

func (s *MemberServiceStruct) UpdateMember(id string, member models.Member) (models.Member, error) {
	existingMember, err := s.memberRepo.GetByID(id)
	if err != nil {
		return models.Member{}, err
	}

	var tagList []models.Tag
	for _, tag := range member.Tags {
		existingTag, err := s.tagRepo.GetByName(tag.Name)
		if err != nil {
			return models.Member{}, err
		}
		if existingTag != nil {
			tagList = append(tagList, *existingTag)
		}
	}
	member.Tags = tagList
	return s.memberRepo.Update(existingMember, member)
}

func (s *MemberServiceStruct) DeleteMember(id string) error {
	return s.memberRepo.Delete(id)
}
