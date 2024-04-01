package models

import "time"

type RefreshToken struct {
	ID           int       `json:"id"`
	RefreshToken string    `json:"refreshToken"`
	UserId       int       `json:"userId"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}
