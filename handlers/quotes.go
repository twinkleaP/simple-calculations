package handlers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQuote(c *gin.Context) {
	quotes := []string{
		"Stay hungry, stay foolish.",
		"Code is like humor. When you have to explain it, it’s bad.",
		"Simplicity is the soul of efficiency.",
		"Before software can be reusable it first has to be usable.",
		"Fix the cause, not the symptom.",
	}
	c.JSON(http.StatusOK, gin.H{
		"quote": quotes[rand.Intn(len(quotes))],
	})
}
