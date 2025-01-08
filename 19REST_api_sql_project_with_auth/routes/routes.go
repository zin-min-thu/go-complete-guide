package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zin-min-thu/apisqlprojectwithauth/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/", home)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// 1 single middleware
	// server.POST("/events", middlewares.Authenticate, createEvent)

	// 2 group middleware
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
