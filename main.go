package main

import (
	"ImageConverter/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":7000")

}
