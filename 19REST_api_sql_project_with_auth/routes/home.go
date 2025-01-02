package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Welcome To My First API Project",
	})
}
