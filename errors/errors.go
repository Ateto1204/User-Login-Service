package errors

import "fmt"

type UserNotFoundError struct {
	Email string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("User with email %s not found", e.Email)
}

type UserExistedError struct {
	Email string
}

func (e *UserExistedError) Error() string {
	return fmt.Sprintf("User with username %s already exists", e.Email)
}

type PwdIncorrectError struct {
	Email string
}

func (e *PwdIncorrectError) Error() string {
	return fmt.Sprintf("Password for user %s is incorrect", e.Email)
}

type EmailInvalidError struct {
	Email string
}

func (e *EmailInvalidError) Error() string {
	return fmt.Sprintf("Email format with %s is invalid", e.Email)
}
