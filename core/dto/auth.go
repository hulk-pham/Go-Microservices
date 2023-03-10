package dto

import "time"

type LoginRequest struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

type SignUpRequest struct {
	FirstName   string    `form:"first_name" json:"first_name" xml:"first_name"  binding:"required"`
	LastName    string    `form:"last_name" json:"last_name" xml:"last_name"  binding:"required"`
	Email       string    `form:"email" json:"email" xml:"email"  binding:"required"`
	Password    string    `form:"password" json:"password" xml:"password" binding:"required"`
	Hobby       string    `form:"hobby" json:"hobby" xml:"hobby"`
	PhoneNumber string    `form:"phone_number" json:"phone_number" xml:"phone_number" binding:"required"`
	Address     string    `form:"address" json:"address" xml:"address"`
	Dob         time.Time `form:"dob" json:"dob" xml:"dob"`
}
