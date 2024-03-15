package routes

import (
	"ImageConverter/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/test", handlers.Welcome)
}
