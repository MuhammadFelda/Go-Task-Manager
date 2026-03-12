package services

import (
	"errors"
	"fmt"

	"github.com/MuhammadFelda/task-manager/internal/models"
)

type SkillCreateRequest struct {
	UserId	uint 	`json:"user_id" binding:"required"`
	Name	string	`json:"name" binding:"required,max=16"`
}

var skills = []models.Skill{
	{Id: 1, UserId: 1, Name: "Golang"},
	{Id: 2, UserId: 1, Name: "PostgreSQL"},
	{Id: 3, UserId: 2, Name: "React"},
}

var skillNextId = 4

type SkillService struct {}

func NewSkillService() *SkillService {
	return &SkillService{}
}

func (s *SkillService) GetAll(userId string) ([]models.Skill, error) {
	if userId == "" {
		return skills, nil
	}

	filtered := []models.Skill{}
	for _, skill := range skills{
		if fmt.Sprintf("%d", skill.UserId) == userId {
			filtered = append(filtered, skill)
		}
	}

	return filtered, nil
}

func (s *SkillService) GetById(id uint) (*models.Skill, error) {
	for _, skill := range skills{
		if skill.Id == id {
			return &skill, nil
		}
	}
	
	return nil, errors.New("Skill Not Found")
}

func (s *SkillService) Create(req SkillCreateRequest) (*models.Skill, error) {
	skill := models.Skill{
		Id: req.UserId,
		Name: req.Name,
	}

	skills = append(skills, skill)
	skillNextId++
	return &skill, nil
}

func (s *SkillService) Delete(id uint) error {
	for index, skill := range skills{
		if skill.Id == id {
			skills =append(skills[:index], skills[index:]...)
			return nil
		}
	}

	return errors.New("Skill Deleted")
}