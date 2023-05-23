package repositories

import (
	"github.com/agusalex/golang-employee-crud/models"
	"github.com/jinzhu/gorm"
	"strings"
)

type SearchRepository struct {
	DB *gorm.DB
}
type SearchRepositoryInterface interface {
	SearchMembers(tags []string, memberType string) ([]models.Member, error)
}

func NewSearchRepository(DB *gorm.DB) *SearchRepository {
	return &SearchRepository{DB: DB}
}
func (r *SearchRepository) SearchMembers(tags []string, memberType string) ([]models.Member, error) {
	var members []models.Member

	interfaceSlice := make([]interface{}, len(tags))
	for i, d := range tags {
		interfaceSlice[i] = d
	}

	query := r.DB.Table("members")
	if len(tags) > 0 {
		placeholderSql := "tags.name IN (" + strings.Repeat("?, ", len(tags)-1) + "?)"
		query = query.Joins("JOIN member_tags ON member_tags.member_id = members.id").
			Joins("JOIN tags ON tags.id = member_tags.tag_id").
			Where(placeholderSql, interfaceSlice...)
	}
	if memberType != "" {
		query = query.Where("members.type = ?", memberType)
	}

	if len(tags) > 0 {
		query = query.Group("members.id").
			Having("COUNT(DISTINCT tags.id) = ?", len(tags))
	}

	if err := query.Find(&members).Error; err != nil {
		return nil, err
	}

	return members, nil
}
