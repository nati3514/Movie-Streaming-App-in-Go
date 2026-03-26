package main

import (

	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
 
	router := gin.Default()

	router.GET("/movies", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Magic Stream Movies Server!",
		})
	})
	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}