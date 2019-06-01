package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"test":    "test",
			"heroku":  "heroku",
			"view":    "view",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
