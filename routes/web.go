package routes

import (
	"auth/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuth(router *gin.Engine) {

	router.GET("/", handlers.Root)
	router.GET("/login", handlers.Login)

}
