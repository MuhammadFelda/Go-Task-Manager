package repository

import (
	"github.com/MuhammadFelda/task-manager/internal/models"
	"gorm.io/gorm"
)

type SkillRepository interface {
	FindAll(userId string) ([]models.Skill, error)
	FindById(id uint) (*models.Skill, error)
	Create(skill *models.Skill) error
	Delete(id uint) error
}

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillRepository{db: db}
}

func (r* skillRepository) FindAll(userId string) ([]models.Skill, error) {
	var skills []models.Skill

	query := r.db.Model(&models.Skill{}).Order("created_at desc")
	if userId != "" {
		query = query.Where("user_id = ?", userId)
	}

	if err := query.Find(&skills).Error; err != nil {
		return nil, err
	}

	return skills, nil
}

func (r *skillRepository) FindById(id uint) (*models.Skill, error) {
	var skill models.Skill
	if err := r.db.First(&skill, id).Error; err != nil {
		return nil, err
	}

	return &skill, nil
}

func (r *skillRepository) Create(skill *models.Skill) error {
	return r.db.Create(skill).Error
}

func (r *skillRepository) Delete(id uint) error {
	return r.db.Delete(&models.Skill{}, id).Error
}