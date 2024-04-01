package models

import "time"

type AccessToken struct {
	ID          int       `json:"id"`
	AccessToken string    `json:"accessToken"`
	UserId      int       `json:"userId"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
