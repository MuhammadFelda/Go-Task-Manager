package services

import (
	"errors"

	"github.com/MuhammadFelda/task-manager/internal/models"
	"github.com/MuhammadFelda/task-manager/internal/repository"
)

type SkillCreateRequest struct {
	UserId	uint 	`json:"user_id" binding:"required"`
	Name	string	`json:"name" binding:"required,max=16"`
}

type SkillService struct {
	repo repository.SkillRepository
}

func NewSkillService(repo repository.SkillRepository) *SkillService {
	return &SkillService{repo: repo}
}

func (s *SkillService) GetAll(userId string) ([]models.Skill, error) {
	return s.repo.FindAll(userId)
}

func (s *SkillService) GetById(id uint) (*models.Skill, error) {
	return s.repo.FindById(id)
}

func (s *SkillService) Create(req SkillCreateRequest) (*models.Skill, error) {
	skill := &models.Skill{
		UserId: req.UserId,
		Name: req.Name,
	}

	if err := s.repo.Create(skill); err != nil {
		return nil, err
	}

	return skill, nil
}

func (s *SkillService) Delete(id uint) error {
	if _, err := s.repo.FindById(id); err != nil {
		return errors.New("Skill Not Found")
	}

	return s.repo.Delete(id)
}