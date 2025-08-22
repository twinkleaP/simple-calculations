package main

import (
	"ayushprasad/Golang/simple-calculations/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register all routes in one place
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
