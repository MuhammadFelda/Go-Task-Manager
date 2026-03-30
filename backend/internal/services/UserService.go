package services

import (
	"errors"

	"github.com/MuhammadFelda/task-manager/internal/models"
	"github.com/MuhammadFelda/task-manager/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserCreateRequest struct {
	Name			string		`json:"name" binding:"required"`
	Email			string		`json:"email" binding:"required"`
	Password		string		`json:"password" binding:"required"`
	Role			string		`json:"role"`
	Avatar			string		`json:"avatar"`
	FaceEmbedding	float64		`json:"face_embedding"`
}

type UserUpdateRequest struct {
	Name			*string		`json:"name"`
	Email			*string		`json:"email"`
	Password		*string		`json:"password"`
	Role			*string		`json:"role"`
	Avatar			*string		`json:"avatar"`
	FaceEmbedding	*float64	`json:"face_embedding"`
}

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetById(id uint) (*models.User, error) {
	return s.repo.FindById(id)
}

func (s *UserService) Create(req UserCreateRequest) (*models.User, error) {
	if _, err := s.repo.FindByEmail(req.Email); err == nil {
		return nil, errors.New("Email Already Taken!")
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Failed to Hash Password")
	}

	role := req.Role
	if role == "" {
		role = "Programmer"
	}

	user := &models.User{
		Name: req.Name,
		Email: req.Email,
		Password: string(passwordHashed),
		Role: role,
		Avatar: req.Avatar,
		FaceEmbedding: req.FaceEmbedding,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Update(id uint, req UserUpdateRequest) (*models.User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New("User Not Found")
	}

	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.Email != nil {
		existingEmail, err := s.repo.FindByEmail(*req.Email)
		if err != nil && existingEmail.ID != id {
			return nil, errors.New("Email Already Taken!")
		}

		user.Email = *req.Email
	}

	if req.Password != nil {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("Failed to Hash Password!")
		}

		user.Password = string(passwordHash)
	}

	if req.Role != nil {
		user.Role = *req.Role
	}

	if req.Avatar != nil {
		user.Avatar = *req.Email
	}

	if req.FaceEmbedding != nil {
		user.FaceEmbedding = *req.FaceEmbedding
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(id uint) error {
	if _, err := s.repo.FindById(id); err != nil {
		return errors.New("User Not Found!")
	}

	return s.repo.Delete(id)
}