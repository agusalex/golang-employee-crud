package repositories

import (
	"github.com/agusalex/golang-employee-crud/models"
	"github.com/jinzhu/gorm"
)

type MemberRepositoryInterface interface {
	GetAll() ([]models.Member, error)
	GetByID(id string) (models.Member, error)
	Save(member models.Member) (models.Member, error)
	Update(existingMember models.Member, member models.Member) (models.Member, error)
	Delete(id string) error
}

type MemberRepository struct {
	DB *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
	return &MemberRepository{
		DB: db,
	}
}

func (r *MemberRepository) GetAll() ([]models.Member, error) {
	var members []models.Member
	if err := r.DB.Preload("Tags").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *MemberRepository) GetByID(id string) (models.Member, error) {
	var member models.Member
	if err := r.DB.Preload("Tags").First(&member, "id = ?", id).Error; err != nil {
		return models.Member{}, err
	}
	return member, nil
}

func (r *MemberRepository) Save(member models.Member) (models.Member, error) {
	if err := r.DB.Save(&member).Error; err != nil {
		return models.Member{}, err
	}
	return member, nil
}

func (r *MemberRepository) Update(existingMember models.Member, newMember models.Member) (models.Member, error) {
	tx := r.DB.Begin()
	if err := tx.Model(&existingMember).Updates(newMember).Error; err != nil {
		tx.Rollback()
		return models.Member{}, err
	}
	if err := tx.Model(&existingMember).Association("Tags").Replace(newMember.Tags).Error; err != nil {
		tx.Rollback()
		return models.Member{}, err
	}
	tx.Commit()
	return newMember, nil
}

func (r *MemberRepository) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.Member{}).Error; err != nil {
		return err
	}
	return nil
}
