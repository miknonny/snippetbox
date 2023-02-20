package models

import "errors"

var (
	ErrNoRecord           = errors.New("models: Not matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

// If the user tries to login with an invalid email and password.
