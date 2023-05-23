package repositories

import (
	"github.com/agusalex/golang-employee-crud/models"
	"github.com/jinzhu/gorm"
)

type TagRepository struct {
	DB *gorm.DB
}
type TagRepositoryInterface interface {
	GetAll() ([]models.Tag, error)
	GetByName(name string) (*models.Tag, error)
}

func NewTagRepository(DB *gorm.DB) *TagRepository {
	return &TagRepository{DB: DB}
}

func (r *TagRepository) GetAll() ([]models.Tag, error) {
	var tags []models.Tag
	if err := r.DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}
func (r *TagRepository) GetByName(name string) (*models.Tag, error) {
	var tag models.Tag
	err := r.DB.Where("name = ?", name).First(&tag).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}
