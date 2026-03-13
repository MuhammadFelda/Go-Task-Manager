package repository

import (
	"github.com/MuhammadFelda/task-manager/internal/models"
	"gorm.io/gorm"
)

type ProjectOwnerRepository interface {
	FindAll(page int) ([]models.ProjectOwner,int64, error)
	FindById(id uint) (*models.ProjectOwner, error)
	Create(projectOwner *models.ProjectOwner) error
	Update(projectOwner *models.ProjectOwner) error
	Delete(id uint) error
}

type projectOwnerRepository struct {
	db *gorm.DB
}

func NewProjectOwnerRepository(db *gorm.DB) ProjectOwnerRepository {
	return &projectOwnerRepository{db: db}
}

func (r *projectOwnerRepository) FindAll(page int) ([]models.ProjectOwner, int64, error) {
	var projectOwner []models.ProjectOwner
	var total int64

	query := r.db.Model(&models.ProjectOwner{}).Where("isDeleted = ?", false)
	query.Count(&total)

	perPage := 10
	offset := (page - 1) * perPage
	err := query.Limit(perPage).Offset(offset).Find(&projectOwner).Error
	if err != nil {
		return nil, 0, err
	}

	return projectOwner, total, nil
}

func (r *projectOwnerRepository) FindById(id uint) (*models.ProjectOwner, error) {
	var projectOwner models.ProjectOwner
	if err := r.db.Where("id = ? AND is_deleted = ?", id, false).First(&projectOwner).Error; err != nil {
		return nil, err
	}

	return &projectOwner, nil
}

func (r *projectOwnerRepository) Create(projectOwner *models.ProjectOwner) error {
	return r.db.Create(projectOwner).Error
}

func (r *projectOwnerRepository) Update(projectOwner *models.ProjectOwner) error {
	return r.db.Save(projectOwner).Error
}

func (r *projectOwnerRepository) Delete(id uint) error {
	return r.db.Model(&models.ProjectOwner{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"is_deleted": true}).Error
}