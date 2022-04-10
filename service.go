package main

import (
	"github.com/gin-gonic/gin"
)

type Product struct {
}

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
