package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
