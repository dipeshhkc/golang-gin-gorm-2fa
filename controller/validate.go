package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Validate -> setup OTP during registration
func Validate(c *gin.Context) {
	if err = otp.Validate(text); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Validation Error"})

	} else {

		c.JSON(http.StatusOK, gin.H{"msg": "Successfully Validated"})
	}
}
