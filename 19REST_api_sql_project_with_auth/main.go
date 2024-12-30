package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zin-min-thu/apisqlprojectwithauth/db"
	"github.com/zin-min-thu/apisqlprojectwithauth/models"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.LoadHTMLGlob("templates/*")

	server.GET("/", home)

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}

func home(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Welcome To My First API Project",
	})
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	// context.JSON(http.StatusOK, gin.H{"message": "Welcome API", "status": true})
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}
