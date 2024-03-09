package models

import customErr "github.com/dafailyasa/golang-template/pkg/custom-errors"

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *AuthRequest) Validate() error {
	if a.Email == "" {
		return customErr.ErrEmailRequired
	}
	if a.Password == "" {
		return customErr.ErrPasswordRequired
	}
	return nil
}
