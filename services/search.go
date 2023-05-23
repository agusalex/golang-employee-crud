package services

import (
	"github.com/agusalex/golang-employee-crud/models"
	"github.com/agusalex/golang-employee-crud/repositories"
)

type SearchServiceStruct struct {
	Repo repositories.SearchRepositoryInterface
}

var SearchService SearchServiceInterface

type SearchServiceInterface interface {
	SearchMembers(tags []string, memberType string) ([]models.Member, error)
}

func NewSearchService(r repositories.SearchRepositoryInterface) SearchServiceInterface {
	return &SearchServiceStruct{Repo: r}
}

func (s *SearchServiceStruct) SearchMembers(tags []string, memberType string) ([]models.Member, error) {
	return s.Repo.SearchMembers(tags, memberType)
}
