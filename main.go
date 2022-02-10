package main

import "github.com/gin-gonic/gin"

func main() {
	//new engine
	r := gin.Default()
	//config route "/"
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!...")
	})
	//run
	// r.Run()
	r.Run(":8000")
}
