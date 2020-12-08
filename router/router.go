package router

import (
	"golang-gin-gorm-2fa/controller"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRoutes : gin router
func SetupRoutes() {
	httpRouter := gin.Default()

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": " API Up and Running"})
	})

	httpRouter.GET("register", controller.Register)
	httpRouter.GET("validate", controller.Validate)

	httpRouter.Run(":5000")

}
