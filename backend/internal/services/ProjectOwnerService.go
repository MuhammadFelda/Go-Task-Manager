package services

import (
	"errors"

	"github.com/MuhammadFelda/task-manager/internal/models"
	"github.com/MuhammadFelda/task-manager/internal/repository"
)

type ProjectOwnerRequest struct {
	Name		string		`json:"name" binding:"required"`
	Creator		uint		`json:"creator"`
}

type ProjectOwnerRequestUpdate struct {
	Name		string		`json:"name" binding:"required"`
	Updater		uint		`json:"updater"`
}

type ProjectOwnerService struct {
	repo repository.ProjectOwnerRepository
}

func NewProjectOwnerService(repo repository.ProjectOwnerRepository) *ProjectOwnerService {
	return &ProjectOwnerService{repo: repo}
}

func (s *ProjectOwnerService) GetAll(page int) ([]models.ProjectOwner,int64, error) {
	return s.repo.FindAll(page)
}

func (s *ProjectOwnerService) Create(req ProjectOwnerRequest) (*models.ProjectOwner, error) {
	projectOwner := &models.ProjectOwner{
		Name: req.Name,
		Creator: req.Creator,
		Updater: req.Creator,
		IsDeleted: false,
	}

	if err := s.repo.Create(projectOwner); err != nil {
		return nil, err
	}

	return projectOwner, nil
}

func (s *ProjectOwnerService) Update(id uint, req ProjectOwnerRequestUpdate) (*models.ProjectOwner, error) {
	projectOwner, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New("Project Owner Not Found")
	}

	projectOwner.Name = req.Name
	projectOwner.Updater = req.Updater

	if err := s.repo.Update(projectOwner); err != nil {
		return nil, err
	}

	return projectOwner, nil
}

func (s *ProjectOwnerService) Delete(id uint) error {
	if _, err := s.repo.FindById(id); err != nil {
		return errors.New("Project Owner Not Found")
	}

	return s.repo.Delete(id)
}