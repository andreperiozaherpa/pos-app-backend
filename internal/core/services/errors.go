package services

import "errors"

var (
	// Auth & User Errors
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrUserInactive       = errors.New("user account is inactive")
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailExists        = errors.New("email already exists")
	ErrUsernameExists     = errors.New("username already exists")
	ErrRoleNotFound       = errors.New("role not found")
	ErrEmployeeNotFound   = errors.New("employee not found")
)
