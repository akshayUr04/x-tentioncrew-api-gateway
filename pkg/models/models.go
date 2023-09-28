package models

type UserDetails struct {
	Name        string `json:"name" validate:"required min=5,max=99"`
	HouseName   string `json:"houseName" validate:"required, min=5,max=30"`
	City        string `json:"city" validate:"required, min=5,max=30"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber int64  `json:"phoneNumber" validate:"required,min=10,max=10"`
}

type Data struct {
	Method int `json:"method"`
	Time   int `json:"time"`
}
