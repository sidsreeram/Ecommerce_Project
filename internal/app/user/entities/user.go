package entities

import (
	"time"

	"gorm.io/gorm"
)

type SignUpRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
	EnteredOTP      string `json:"otp"`
}

type SignUpResponse struct {
	Message string `json:"message"`
}

type OTPToken struct {
	gorm.Model
	Email          string // Replace with the appropriate user reference type
	OTPCode        string `gorm:"size:6"`
	ExpirationTime time.Time
}
