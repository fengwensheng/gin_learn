package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title       string
	Description string
	Content     string
}

func main() {
	//new engine
	r := gin.Default()
	//for return type: HTML
	r.LoadHTMLGlob("templates/*.html")
	//config route "/"
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!...")
	})
	//restful api, CURD
	r.GET("/get", func(c *gin.Context) {
		//Browser can only test this kind of http method
		c.String(http.StatusOK, "Get sf Retrieve")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "Post sf Create")
	})
	r.PUT("/put", func(c *gin.Context) {
		c.String(http.StatusOK, "Put sf Update")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "Delete sf Delete")
	})
	//return type
	//json
	r.GET("/json", func(c *gin.Context) {
		// c.JSON(http.StatusOK, map[string]interface{}{
		// 	"id":   1,
		// 	"name": "Vincent Feng",
		// })

		// c.JSON(http.StatusOK, gin.H{
		// 	"id":   1,
		// 	"name": "Vincent Feng",
		// })

		article := Article{
			Title:       "title",
			Description: "description",
			Content:     "content",
		}
		c.JSON(http.StatusOK, &article)
	})
	//jsonp
	//for cross-domain
	//input:(...?callback=xxx)
	// http://localhost:8000/jsonp?callback=xxx
	//output:(xxx(json))
	// xxx({
	// 	"Title": "title",
	// 	"Description": "description",
	// 	"Content": "content"
	//   });
	r.GET("/jsonp", func(c *gin.Context) {
		article := Article{
			Title:       "title",
			Description: "description",
			Content:     "content",
		}
		c.JSONP(http.StatusOK, &article)
	})
	//xml
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, &gin.H{
			"success": true,
			"msg":     "xml content",
		})
	})
	//html
	//SSR
	//pre mkdir templates
	r.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", &gin.H{
			"title": "SERVER NEWS",
		})
	})
	//run
	// r.Run()
	r.Run(":8000")
}
