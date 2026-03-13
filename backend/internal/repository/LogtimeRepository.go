package repository

import (
	"github.com/MuhammadFelda/task-manager/internal/models"
	"gorm.io/gorm"
)

type LogtimeRepository interface {
	FindAll(userId string, from string, to string, page int) ([]models.Logtime, int64, error)
	FindById(id uint) (*models.Logtime, error)
	Create(logtime *models.Logtime) error
	Delete(id uint) error
}

type logtimeRepository struct {
	db *gorm.DB
}

func NewLogtimeRepository(db *gorm.DB) LogtimeRepository {
	return &logtimeRepository{db: db}
}

func (r *logtimeRepository) FindAll(userId string, from string, to string, page int) ([]models.Logtime, int64, error) {
	var logtimes []models.Logtime
	var total int64

	query := r.db.Model(&models.Logtime{}).Preload("Task").Preload("User")

	if userId != "" {
		query = query.Where("user_id = ?", userId)
	}

	if from != "" {
		query = query.Where("date BETWEEN ? AND ?", from, to)
	}

	query.Count(&total)

	perPage := 10
	offset := (page - 1) * perPage

	err := query.Order("date desc").Limit(perPage).Offset(offset).Find(&logtimes).Error
	if err != nil {
		return  nil, 0, err
	}

	return logtimes, total, nil
}

func (r *logtimeRepository) FindById(id uint) (*models.Logtime, error) {
	var logtime models.Logtime
	if err := r.db.Preload("Task").Preload("User").First(&logtime, id).Error; err != nil {
		return nil, err
	}

	return &logtime, nil
}

func (r *logtimeRepository) Create(logtime *models.Logtime) error {
	return r.db.Create(logtime).Error
}

func (r *logtimeRepository) Delete(id uint) error {
	return r.db.Delete(&models.Logtime{}, id).Error
}