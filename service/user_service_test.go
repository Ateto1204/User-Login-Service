package service_test

import (
	"testing"
	"user-app/errors"
	"user-app/repository"
	"user-app/service"
)

func TestUserService(t *testing.T) {
	repo := repository.NewUserRepository()
	service := service.NewUserService(repo)

	// Test user not found error
	_, err := service.GetUserByEmail("nonexisted_user", "pwd")
	if _, ok := err.(*errors.UserNotFoundError); !ok {
		t.Errorf("Expected UserNotFoundError, got %v", err)
	}

	// Test user existed error
	_ = service.CreateUser("existing_user", "existing@gmail.com", "existing")
	err = service.CreateUser("existing_user", "existing@gmail.com", "existing")
	if _, ok := err.(*errors.UserExistedError); !ok {
		t.Errorf("Expected UserExistedError, got %v", err)
	}

	// Test pwd incorrect error
	_, err = service.GetUserByEmail("existing@gmail.com", "incorrect_pwd")
	if _, ok := err.(*errors.PwdIncorrectError); !ok {
		t.Errorf("Expected PwdIncorrectError, got %v", err)
	}
}
