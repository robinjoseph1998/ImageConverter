package routes

import (
	"ImageConverter/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/pngconvert", handlers.JpgToPngConverter)
	r.POST("/gifconvert", handlers.JpgToGifConverter)
	r.POST("/jpegconvert", handlers.JpgToJpegConverter)
}
