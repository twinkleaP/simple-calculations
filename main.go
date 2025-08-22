package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Route: Addition
	r.GET("/add", func(c *gin.Context) {
		aStr := c.Query("a")
		bStr := c.Query("b")

		a, err1 := strconv.Atoi(aStr)
		b, err2 := strconv.Atoi(bStr)

		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, please provide two integers."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": a + b})
	})

	// Route: substraction
	r.GET("/add", func(c *gin.Context) {
		aStr := c.Query("a")
		bStr := c.Query("b")

		a, err1 := strconv.Atoi(aStr)
		b, err2 := strconv.Atoi(bStr)

		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, please provide two integers."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": a - b})
	})

	// Route: multiplication
	r.GET("/add", func(c *gin.Context) {
		aStr := c.Query("a")
		bStr := c.Query("b")

		a, err1 := strconv.Atoi(aStr)
		b, err2 := strconv.Atoi(bStr)

		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, please provide two integers."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": a * b})
	})
	// Start server
	r.Run(":8080")
}
