package models

import "time"

type Address struct {
	ID             int       `json:"id"`
	UserId         int       `json:"userId"`
	Header         string    `json:"header"`
	Name           string    `json:"name"`
	Country        string    `json:"country"`
	CityId         int       `json:"cityId"`
	DistrictId     int       `json:"districtId"`
	NeighborhoodId int       `json:"neighborhoodId"`
	Description    string    `json:"description"`
	Phone          string    `json:"phone"`
	BillType       bool      `json:"billType"`
	SSN            string    `json:"ssn"`
	IsPrimary      bool      `json:"isPrimary"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	City           string    `json:"city"`
	District       string    `json:"district"`
	Neighborhood   string    `json:"neighborhood"`
}
