package main

import (

	"github.com/gin-gonic/gin"
	"fmt"
	controller "github.com/nati3514/Movie-Streaming-App-in-Go/Server/MagicStreamMoviesServer/controllers"
)

func main() {
 
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Magic Stream Movies Server!",
		})
	})

	router.GET("/movies", controller.GetMovies())
	router.GET("/movies/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())

	if err := router.Run(":3005"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}