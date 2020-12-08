package controller

import (
	"bytes"
	"crypto"
	"golang-gin-gorm-2fa/database"
	"golang-gin-gorm-2fa/model"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sec51/twofactor"
)

var (
	email  = "dipesh.kc@wesionary.team"
	issuer = "wesionaryTEAM"
)

//Register -> setup OTP during registration
func Register(c *gin.Context) {
	// Instantiate totp
	otp, err := twofactor.NewTOTP(email, issuer, crypto.SHA1, 6)
	if err != nil {
		log.Print("Error during OTP generation", err)
	}
	db := database.DBConnection()

	otpToByte, _ := otp.ToBytes()

	ToFactorRecord := model.TwoFactor{OTP: otpToByte, Email: email}

	if db.Create(&ToFactorRecord).Error != nil {
		log.Panic("Unable to add new otp byte", err)
	}

	// Generate QR bytes
	qrBytes, err := otp.QR()
	if err != nil {
		log.Print("Error during QR generation", err)
	}
	//Generate PNG from QR bytes
	img, _, _ := image.Decode(bytes.NewReader(qrBytes))
	out, err := os.Create("./QRImg.png")
	err = png.Encode(out, img)

	c.JSON(http.StatusOK, gin.H{"msg": "Successfully Registered", "data": "link to QR"})
}
