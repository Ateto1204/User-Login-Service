package service_test

import (
	"log"
	"testing"

	"github.com/Ateto/User-Login-Service/db"
	"github.com/Ateto/User-Login-Service/errors"
	"github.com/Ateto/User-Login-Service/repository"
	"github.com/Ateto/User-Login-Service/service"
)

func TestUserService(t *testing.T) {
	db, err := db.NewDB("../config.json", "../db/init.sql")
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)

	// Test invalid email format
	err = service.CreateUser("invalid", "invalid", "invalid")
	if _, ok := err.(*errors.EmailInvalidError); !ok {
		t.Errorf("Expected EmailInvalidError, got %v", err)
	}

	_, err = service.GetUserByEmail("invalid", "invalid")
	if _, ok := err.(*errors.EmailInvalidError); !ok {
		t.Errorf("Expected EmailInvalidError, got %v", err)
	}

	// Test user not found error
	_, err = service.GetUserByEmail("nonexisted@gmail.com", "nonexisted")
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
