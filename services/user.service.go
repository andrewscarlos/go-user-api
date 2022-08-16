package services

import (
	"go-web/models"
	"go-web/repository"
)

type UserServiceInterface interface {
	Find(id string) (user *models.User, err error)
	FindAll() (user []*models.User, err error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id string) error
}

type userService struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserService(userRepository repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{userRepository}
}

func (s *userService) Find(id string) (user *models.User, err error) {
	return s.userRepository.Find(id)
}

func (s *userService) FindAll() (user []*models.User, err error) {
	return s.userRepository.FindAll()
}

func (s *userService) Create(user *models.User) error {
	return s.userRepository.Create(user)
}

func (s *userService) Update(user *models.User) error {
	return s.userRepository.Update(user)
}

func (s *userService) Delete(id string) error {
	return s.userRepository.Delete(id)
}
