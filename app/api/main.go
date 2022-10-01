package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/quikzens/rest-api-boilerplate/config"
	"github.com/quikzens/rest-api-boilerplate/routes"
)

func main() {
	if config.EnvMode == config.ProductionMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-type"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))

	routes.Set(r)

	log.Fatal(r.Run(config.ServerAddress))
}
