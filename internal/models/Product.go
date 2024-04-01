package models

type Product struct {
	ID      int                    `json:"id"`
	UserId  int                    `json:"userId"`
	Product map[string]interface{} `json:"product"`
}
