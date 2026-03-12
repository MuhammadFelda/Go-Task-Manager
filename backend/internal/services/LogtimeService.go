package services

import (
	"errors"
	"time"

	"github.com/MuhammadFelda/task-manager/internal/models"
)

type LogtimeCreateRequest struct {
	UserID      uint    `json:"user_id" binding:"required"`
	TaskId      uint    `json:"task_id" binding:"required"`
	Date        string  `json:"date" binding:"required"`
	TimeUsed    float64 `json:"time_used" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

// DUMMY DATA
var logtimes = []models.Logtime{
	{Id: 1, UserId: 1, TaskId: 1, Date: time.Now(), TimeUsed: 3.5, Description: "fix login bug"},
	{Id: 2, UserId: 1, TaskId: 2, Date: time.Now(), TimeUsed: 2.0, Description: "code review"},
}

var logtimeNextId = 4

type LogtimeService struct{}

func NewLogtimeService() *LogtimeService {
	return &LogtimeService{}
}

func (s *LogtimeService) GetAll(userId string, from string, to string, page int) ([]models.Logtime, error) {
	return logtimes, nil
}

func (s *LogtimeService) GetById(id uint) (*models.Logtime, error) {
	for _, logtime := range logtimes {
		if logtime.Id == id {
			return &logtime, nil
		}
	}

	return nil, errors.New("Logtimes Not Found")
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

	logtime := models.Logtime{
		Id: uint(logtimeNextId),
		UserId: req.UserID,
		TaskId: req.TaskId,
		Date: date,
		TimeUsed: timeUsed,
		Description: req.Description,
	}

	logtimes = append(logtimes, logtime)
	logtimeNextId++

	return &logtime, nil
}

func (s *LogtimeService) Delete(id uint) error {
	for i, logtime := range logtimes {
		if logtime.Id == id {
			logtimes = append(logtimes[:i], logtimes[i+1:]...)
			return nil
		}
	}

	return errors.New("Logtime Not Found")
}