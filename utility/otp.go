package utility

import (
	"golang-gin-gorm-2fa/database"
	"golang-gin-gorm-2fa/model"
	"log"

	"github.com/sec51/twofactor"
)

var (
	issuer = "wesionaryTEAM"
)

func getOTPFromDB() *twofactor.Totp {
	db := database.DBConnection()
	var record1 model.TwoFactor
	if err := db.First(&record1).Error; err != nil {
		log.Print("Error during retrieving from DB: ", err)
	}

	otp, err := twofactor.TOTPFromBytes(record1.OTP, issuer)
	if err != nil {
		log.Print("Error during converting bytes to TOTP: ", err)
	}
	return otp
}
