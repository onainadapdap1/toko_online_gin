package service

import (
	"errors"
	"toko_online_gin/helpers"
	"toko_online_gin/models"
	"toko_online_gin/repository"
)

type UserServiceInterface interface {
	RegisterUser(input *models.RegisterUserInput) (*models.User, error)
	LoginUser(input models.LoginUserInput) (models.User, error)
	IsEmailAvailable(input models.CheckEmailInput) (bool, error)
	SaveAvatar(ID uint, fileLocation string) (models.User, error)
	GetUserByID(ID uint) (models.User, error)
}

type userService struct {
	repo repository.UserRepoInterface
}

func NewUserService(repo repository.UserRepoInterface) UserServiceInterface {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(input *models.RegisterUserInput) (*models.User, error) {
	// set new user
	user := models.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
		Role: "user",
	}

	newUser, err := s.repo.RegisterUser(&user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) LoginUser(input models.LoginUserInput) (models.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	comparePass := helpers.ComparePassword([]byte(user.Password), []byte(password))
	if !comparePass {
		return user, err
	}

	return user, nil
}

func (s *userService) IsEmailAvailable(input models.CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID != 0 {
		return false, nil
	}

	return true, nil
}

func (s *userService) SaveAvatar(ID uint, fileLocation string) (models.User, error) {
	user, err := s.repo.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation
	updatedUser, err := s.repo.SaveAvatar(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *userService) GetUserByID(ID uint) (models.User, error) {
	var user models.User

	user, err := s.repo.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}