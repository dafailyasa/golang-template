package customErr

import "errors"

var (
	ErrInvalidToken = errors.New("token was invalid")
	// ErrExpiredToken is returned when access token was expired
	ErrExpiredToken = errors.New("token was expired")
	// ErrMongoUriRequired is returned when a mongoUri is not provided
	ErrMongoUrlRequired = errors.New("MONGO_URI is required")
)
