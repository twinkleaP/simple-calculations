package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMath(c *gin.Context) {
	aStr := c.DefaultQuery("a", "")
	bStr := c.DefaultQuery("b", "")
	op := c.DefaultQuery("op", "random")

	var a, b int
	var err error

	if aStr == "" || bStr == "" {
		rand.Seed(time.Now().UnixNano())
		a = rand.Intn(20)
		b = rand.Intn(20) + 1
	} else {
		a, err = strconv.Atoi(aStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid 'a' parameter"})
			return
		}
		b, err = strconv.Atoi(bStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid 'b' parameter"})
			return
		}
	}

	ops := []string{"add", "subtract", "multiply", "divide"}
	if op == "random" {
		op = ops[rand.Intn(len(ops))]
	}

	var result float64
	switch op {
	case "add":
		result = float64(a + b)
	case "subtract":
		result = float64(a - b)
	case "multiply":
		result = float64(a * b)
	case "divide":
		if b == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "division by zero"})
			return
		}
		result = float64(a) / float64(b)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid operation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"a":         a,
		"b":         b,
		"operation": op,
		"result":    result,
	})
}
