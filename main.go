package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//new engine
	r := gin.Default()
	//config route "/"
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!...")
	})
	//restful api, CURD
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "Get sf Retrieve")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "Post sf Create")
	})
	r.PUT("/put", func(c *gin.Context) {
		c.String(http.StatusOK, "Put sf Update")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "Delete sf delete")
	})
	//run
	// r.Run()
	r.Run(":8000")
}
