package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zin-min-thu/apisqlprojectwithauth/db"
	"github.com/zin-min-thu/apisqlprojectwithauth/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.LoadHTMLGlob("templates/*")

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
