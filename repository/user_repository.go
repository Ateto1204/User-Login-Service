package repository

import (
	"database/sql"
	"log"

	"github.com/Ateto/User-Login-Service/db"
	"github.com/Ateto/User-Login-Service/errors"
	"github.com/Ateto/User-Login-Service/model"
)

type UserRepository struct {
	db *db.MysqlDatabase
}

func NewUserRepository(db *db.MysqlDatabase) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) GetUserByEmail(email, pwd string) (*model.User, error) {
	err := isUserExisted(repo.db.DB, email)
	if _, ok := err.(*errors.UserNotFoundError); ok {
		return nil, err
	}

	query := `SELECT email, name, pwd FROM user WHERE email = ?`
	row := repo.db.DB.QueryRow(query, email)

	var user model.User
	err = row.Scan(&user.Email, &user.Name, &user.Pwd)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &errors.UserNotFoundError{Email: email}
		}
		return nil, err
	}

	if pwd != user.Pwd {
		return nil, &errors.PwdIncorrectError{Email: email}
	}
	return &user, nil
}

func (repo *UserRepository) CreateUser(user *model.User) error {
	email := user.Email
	err := isUserExisted(repo.db.DB, email)
	if _, ok := err.(*errors.UserExistedError); ok {
		return err
	}
	query := `INSERT INTO user (email, name, pwd) VALUES (?, ?, ?)`
	_, err = repo.db.DB.Exec(query, user.Email, user.Name, user.Pwd)
	if err != nil {
		return err
	}
	return nil
}

func isUserExisted(db *sql.DB, email string) error {
	query := `SELECT email FROM user WHERE email = ?`
	row := db.QueryRow(query, email)

	var existingEmail string
	err := row.Scan(&existingEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return &errors.UserNotFoundError{Email: email}
		}
		log.Fatal(err)
	}
	return &errors.UserExistedError{Email: email}
}
