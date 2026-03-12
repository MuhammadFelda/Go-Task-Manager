package services

import (
	"errors"

	"github.com/MuhammadFelda/task-manager/internal/models"
)

type ProjectOwnerRequest struct {
	Name		string		`json:"name" binding:"required"`
	Creator		uint		`json:"creator"`
}

type ProjectOwnerRequestUpdate struct {
	Name		string		`json:"name" binding:"required"`
	Updater		uint		`json:"updater"`
}

var projectOwners = []models.ProjectOwner{
	{Id: 1, Name: "KNDI", Creator: 1, Updater: 1, IsDeleted: false},
	{Id: 2, Name: "KD", Creator: 1, Updater: 1, IsDeleted: false},
	{Id: 3, Name: "External", Creator: 1, Updater: 1, IsDeleted: false},
}

var nextProjectownerId uint = 4

type ProjectOwnerService struct {}

func NewProjectOwnerService() *ProjectOwnerService {
	return &ProjectOwnerService{}
}

func (s *ProjectOwnerService) GetAll(page int) ([]models.ProjectOwner, error) {
	active := []models.ProjectOwner{}
	for _, projectOwner := range projectOwners{
		if !projectOwner.IsDeleted {
			active = append(active, projectOwner)
		}
	}
	return active, nil
}

func (s *ProjectOwnerService) Create(req ProjectOwnerRequest) (*models.ProjectOwner, error) {
	projectOwner := models.ProjectOwner{
		Id: nextProjectownerId,
		Name: req.Name,
		Creator: req.Creator,
		Updater: req.Creator,
		IsDeleted: false,
	}

	projectOwners = append(projectOwners, projectOwner)
	nextProjectownerId++

	return &projectOwner, nil
}

func (s *ProjectOwnerService) Update(id uint, req ProjectOwnerRequestUpdate) (*models.ProjectOwner, error) {
	for index, projectOwner := range projectOwners{
		if projectOwner.Id == id && !projectOwner.IsDeleted {
			projectOwners[index].Name = req.Name
			projectOwners[index].Updater = req.Updater
			return &projectOwners[index], nil
		}
	}

	return nil, errors.New("Project Owner Not Found")
}

func (s *ProjectOwnerService) Delete(id uint) error {
	for index, projectOwner := range projectOwners{
		if projectOwner.Id == id && !projectOwner.IsDeleted {
			projectOwners[index].IsDeleted = true
			return nil
		}
	}
	
	return errors.New("Project Owner Not Found")
}