package models

import "time"

type ProductComment struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Star        int       `json:"star"`
	UserId      int       `json:"userId"`
	ProductId   int       `json:"productId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
