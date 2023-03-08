package user

import (
	"hulk/go-webservice/common"
	"time"
)

type User struct {
	common.Model
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Hobby       string    `json:"hobby"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Dob         time.Time `json:"dob"`
	Avatar      string    `json:"avatar"`
}
