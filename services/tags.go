package services

import (
	"github.com/agusalex/golang-employee-crud/models"
	"github.com/agusalex/golang-employee-crud/repositories"
)

type TagServiceInterface interface {
	GetAllTags() ([]models.Tag, error)
	GetTagByName(name string) (*models.Tag, error)
}

var TagService TagServiceInterface

type TagsServiceStruct struct {
	Repo repositories.TagRepositoryInterface
}

func NewTagService(r repositories.TagRepositoryInterface) TagServiceInterface {
	return &TagsServiceStruct{Repo: r}
}

func (s *TagsServiceStruct) GetAllTags() ([]models.Tag, error) {
	return s.Repo.GetAll()
}

func (s *TagsServiceStruct) GetTagByName(name string) (*models.Tag, error) {
	return s.Repo.GetByName(name)
}
