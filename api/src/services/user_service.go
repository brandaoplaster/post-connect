package services

import (
	"api/api/src/models"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	All() ([]models.User, error)
	FindById(id uint64) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(id uint64, user models.User) (models.User, error)
	Delete(id uint64) error
}

type UserService interface {
	List() ([]models.User, error)
	Find(id uint64) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(id uint64, user models.User) (models.User, error)
	Delete(id uint64) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) List() ([]models.User, error) {
	return s.repo.All()
}

func (s *userService) Find(id uint64) (models.User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return models.User{}, err
	}
	if user.ID == 0 {
		return models.User{}, ErrUserNotFound
	}
	return user, nil
}

func (s *userService) Create(user models.User) (models.User, error) {
	if err := user.Prepare("Create"); err != nil {
		return models.User{}, err
	}
	return s.repo.Create(user)
}

func (s *userService) Update(id uint64, user models.User) (models.User, error) {
	if err := user.Prepare("update"); err != nil {
		return models.User{}, err
	}
	user, err := s.repo.Update(id, user)
	if err != nil {
		if err.Error() == "user not found" {
			return models.User{}, ErrUserNotFound
		}
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) Delete(id uint64) error {
	err := s.repo.Delete(id)
	if err != nil && err.Error() == "user not found" {
		return ErrUserNotFound
	}
	return err
}
