package errors

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordIncorrect    = errors.New("password is incorrect")
	ErrUsernameExist        = errors.New("username already exist")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUsernameExist,
	ErrPasswordDoesNotMatch,
}
