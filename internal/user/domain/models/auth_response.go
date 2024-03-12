package models

import customErr "github.com/dafailyasa/golang-template/pkg/custom-errors"

type AuthResponse struct {
	Token string `json:"token"`
}

func (a *AuthResponse) Validate() error {
	if a.Token == "" {
		return customErr.ErrExpiredToken
	}
	return nil
}
