package service

import (
	"regexp"

	"github.com/Ateto/User-Login-Service/errors"
	"github.com/Ateto/User-Login-Service/model"
	"github.com/Ateto/User-Login-Service/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repository: repo}
}

func (service *UserService) GetUserByEmail(email, pwd string) (*model.User, error) {
	err := CheckEmailFormat(email)
	if err != nil {
		return nil, err
	}
	return service.repository.GetUserByEmail(email, pwd)
}

func (service *UserService) CreateUser(name, email, pwd string) error {
	err := CheckEmailFormat(email)
	if err != nil {
		return err
	}
	user := &model.User{
		Name:  name,
		Email: email,
		Pwd:   pwd,
	}
	return service.repository.CreateUser(user)
}

func CheckEmailFormat(email string) error {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegexPattern)
	isValid := re.MatchString(email)
	if isValid {
		return nil
	}
	return &errors.EmailInvalidError{Email: email}
}
