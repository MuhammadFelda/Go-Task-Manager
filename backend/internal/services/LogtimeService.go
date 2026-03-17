package services

import (
	"errors"
	"time"

	"github.com/MuhammadFelda/task-manager/internal/models"
	"github.com/MuhammadFelda/task-manager/internal/repository"
)

type LogtimeCreateRequest struct {
	UserID      uint    `json:"user_id" binding:"required"`
	TaskId      uint    `json:"task_id" binding:"required"`
	Date        string  `json:"date" binding:"required"`
	TimeUsed    float64 `json:"time_used" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

type LogtimeService struct{
	repo repository.LogtimeRepository
}

func NewLogtimeService(repo repository.LogtimeRepository) *LogtimeService {
	return &LogtimeService{repo: repo}
}

func (s *LogtimeService) GetAll(userId string, from string, to string, page int) ([]models.Logtime, int64, error) {
	return s.repo.FindAll(userId, from, to, page)
}

func (s *LogtimeService) GetById(id uint) (*models.Logtime, error) {
	return s.repo.FindById(id)
}

func (s *LogtimeService) Create(req LogtimeCreateRequest) (*models.Logtime, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("Invalid Date Format")
	}

	timeUsed := req.TimeUsed
	if timeUsed < 0 {
		timeUsed = 1
	}

	logtime := &models.Logtime{
		UserId: req.UserID,
		TaskId: req.TaskId,
		Date: date,
		TimeUsed: timeUsed,
		Description: req.Description,
	}

	if err := s.repo.Create(logtime); err != nil {
		return nil, err
	}

	return logtime, nil
}

func (s *LogtimeService) Delete(id uint) error {
	if _, err := s.repo.FindById(id); err != nil {
		return errors.New("Logtime Not Found")
	}

	return s.repo.Delete(id)
}