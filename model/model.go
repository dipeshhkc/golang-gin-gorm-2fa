package model

import "gorm.io/gorm"

// TwoFactor is the model for the TwoFactor table.
type TwoFactor struct {
	gorm.Model
	Email string `gorm:"unique"`
	OTP   []byte `gorm:"not null"`
}
