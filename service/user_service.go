package service

import (
	"user-app/model"
	"user-app/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repository: repo}
}

func (service *UserService) GetUserByEmail(email, pwd string) (*model.User, error) {
	return service.repository.GetUserByEmail(email, pwd)
}

func (service *UserService) CreateUser(name, email, pwd string) error {
	user := &model.User{
		Name:  name,
		Email: email,
		Pwd:   pwd,
	}
	return service.repository.CreateUser(user)
}
