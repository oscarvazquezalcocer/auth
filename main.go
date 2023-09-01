package main

import (
	"auth/config"
	"auth/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.ConfigureViper()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	routes.SetupAuth(r)

	r.Run(viper.GetString("PORT"))
}
