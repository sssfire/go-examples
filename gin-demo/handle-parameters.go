package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		fmt.Println(action)
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.GET("/user1/:name/*action", func(c *gin.Context) {
		fullPath := c.FullPath()
		fmt.Println("the full path is %s", fullPath)
		c.String(http.StatusOK, fullPath)
		// c.FullPath() == "/user/:name/*action" // true
	})

	// 解析url中的参数
	// http://localhost:8080/welcome?firstname="aaa"&&lastname="bbb"
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":8080")
}
