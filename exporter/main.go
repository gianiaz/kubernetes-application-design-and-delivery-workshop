package main

import (
	"github.com/gin-gonic/gin"
	"math"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/load", load)
	r.Run(":3000")
}

func load(c *gin.Context) {
	now := float64(time.Now().Unix())
	c.JSON(200, gin.H{"load": math.Abs(math.Cos(now / 100))})
}
