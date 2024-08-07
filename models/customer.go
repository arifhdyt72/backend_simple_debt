package models

type Customer struct {
	MavisModel
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Status      *bool  `json:"status"`
}
