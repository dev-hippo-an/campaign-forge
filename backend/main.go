package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", IndexHanlder())

	router.Run(":8080")
}
