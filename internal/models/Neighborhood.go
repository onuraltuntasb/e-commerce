package models

type Neighborhood struct {
	ID         int    `json:"id"`
	DistrictId int    `json:"districtId"`
	Name       string `json:"name"`
	PostalCode string `json:"postalCode"`
}
