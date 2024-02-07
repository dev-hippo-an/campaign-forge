package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHanlder() func(*gin.Context) {

	return func(ctx *gin.Context) {
		response := GetMessageFromOpenAI()
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": *response,
		})
	}

}
