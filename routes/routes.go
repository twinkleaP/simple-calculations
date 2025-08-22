package routes

import (
	"ayushprasad/Golang/simple-calculations/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/quotes", handlers.GetQuote)
	r.GET("/math", handlers.GetMath)
}
