package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(g *gin.Context){
			g.JSON(200, gin.H{
				"message": "Hello World!",
			})
		})
	}

	router.Run(":5000")
}