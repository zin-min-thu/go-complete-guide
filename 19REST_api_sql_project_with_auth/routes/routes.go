package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/", home)

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEvent)
	server.PUT("/events/:id", updateEvent)
}
