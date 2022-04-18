package models

import (
	"time"
)

type User struct {
	UserId   uint64    `json:"user_id,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Fullname string    `json:"fullname,omitempty"`
	Dob      time.Time `json:"dob,omitempty"`
}
