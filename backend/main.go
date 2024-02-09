package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/campaigns", CampaignListCreateHandler())
	router.POST("/campaign", CampaignCreateHandler())

	log.Fatal(router.Run(":8080"))
}
