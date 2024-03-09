package customErr

import "errors"

var (
	// ErrPasswordConfirmation is returned when a password confirmation is not same as password
	ErrPasswordConfirmation = errors.New("password confirmation is not same as password")
	// ErrEmailRequired is returned when an email address is not provided
	ErrEmailRequired = errors.New("email address is required")
	// ErrEmailRegistered is returned when an email address is not provided
	ErrEmailRegistered = errors.New("email address already register")
	// ErrPasswordRequired is returned when a password is not provided
	ErrPasswordRequired = errors.New("password is required")
	// ErrInvalidToken is returned when a token is not provided
	ErrInvalidToken = errors.New("token was invalid")
	// ErrTokenRequired is returned when a token is not provided
	ErrTokenRequired = errors.New("token is required")
	// ErrExpiredToken is returned when access token was expired
	ErrExpiredToken = errors.New("token was expired")
	// ErrMongoUriRequired is returned when a mongoUri is not provided
	ErrMongoUrlRequired = errors.New("MONGO_URI is required")
	// ErrInvalidEmail is returned when an email is not valid
	ErrInvalidEmail = errors.New("email is not valid")
	// ErrPasswordLength is returned when a password is not valid
	ErrPasswordLength = errors.New("password must be at least 8 characters")
	// ErrPasswordFormat is returned when a password is not valid
	ErrPasswordFormat = errors.New("password must contain at least one uppercase letter, one lowercase letter, one number and one special character")
	// ErrCreateAtRequired is returned when a create_at is not provided
	ErrCreateAtRequired = errors.New("create_at is required")
	// ErrInternalServer is returned when a server error is not provided
	ErrInternalServer = errors.New("internal server error")
)
