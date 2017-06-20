package main

import (
	"flag"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

  r.Use(gin.Logger())
  r.Use(gin.Recovery())

	flag.Parse()
	r.GET("/ws", stream)
	log.Printf("serving on port 8080")
	r.Run(":8080")
}
