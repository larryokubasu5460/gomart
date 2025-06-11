package service

import (
	"errors"
	"github.com/larryokubasu5460/gomart/user-service/models"
	"github.com/larryokubasu5460/gomart/user-service/repository"
	"github.com/larryokubasu5460/gomart/user-service/utils"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) Register(username, email, password string) error {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Email: email,
		Password: hashed,
	}

	return s.Repo.Create(user)
}

func (s *UserService) Login(email, password string)(string,error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if !utils.CheckPasswordHash(password,user.Password) {
		return "", errors.New("invalid credentials")
	}
	token, err := utils.GenerateJWT(user.ID, user.Email)
	return token, err
}