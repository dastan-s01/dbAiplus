package models

type Employee struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	City    string `json:"city"`
	Phone   string `json:"phone"`
}
