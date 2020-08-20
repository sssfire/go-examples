package main

import "github.com/gin-gonic/gin"

//install gin first
//$go get -u github.com/gin-gonic/gin
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
