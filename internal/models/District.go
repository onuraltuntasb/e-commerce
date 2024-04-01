package models

type District struct {
	ID     int    `json:"id"`
	CityId int    `json:"cityId"`
	Name   string `json:"name"`
}
