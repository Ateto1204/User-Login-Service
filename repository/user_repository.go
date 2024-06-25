package repository

import (
	"user-app/errors"
	"user-app/model"
)

type UserRepository struct {
	users map[string]*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[string]*model.User)}
}

func (repo *UserRepository) GetUserByEmail(email, pwd string) (*model.User, error) {
	user := repo.users[email]
	if user == nil {
		return nil, &errors.UserNotFoundError{Email: email}
	}
	if pwd != user.Pwd {
		return nil, &errors.PwdIncorrectError{Email: email}
	}
	return user, nil
}

func (repo *UserRepository) CreateUser(user *model.User) error {
	email := user.Email
	if repo.users[email] != nil {
		return &errors.UserExistedError{Email: email}
	}
	repo.users[email] = user
	return nil
}
