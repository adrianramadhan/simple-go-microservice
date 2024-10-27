package service

import (
	"time"
	"user-service/internal/dto"
	"user-service/internal/entity"
	"user-service/internal/repository"
)

type UserService interface {
	CreateUser(dto *dto.CreateUserRequest) (dto.CreateUserResponse, error)
	GetAllUsers() ([]dto.CreateUserResponse, error)
	GetUserById(id uint) (dto.CreateUserResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository}
}

func (s *userService) CreateUser(req *dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	user := entity.User{
		Email:     req.Email,
		Password:  req.Password,
		Username:  req.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepository.Create(&user); err != nil {
		return dto.CreateUserResponse{}, err
	}

	return dto.CreateUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) GetAllUsers() ([]dto.CreateUserResponse, error) {
	users, err := s.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var response []dto.CreateUserResponse
	for _, user := range users {
		response = append(response, dto.CreateUserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return response, nil
}

func (s *userService) GetUserById(id uint) (dto.CreateUserResponse, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	return dto.CreateUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
