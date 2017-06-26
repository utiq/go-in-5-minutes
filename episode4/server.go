package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.SetHTMLTemplate(template.Must(template.ParseFiles("index.html")))
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	h := newHub()
	wsh := wsHandler{h: h}
	r.GET("/ws", func(c *gin.Context) {
		wsh.ServeHTTP(c.Writer, c.Request)
	})
	log.Printf("serving on port 8080")
	r.Run(":8080")
}
