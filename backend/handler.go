package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func CampaignListCreateHandler() func(*gin.Context) {
	return func(ctx *gin.Context) {

		campaignCategory := CampaignCategory{}

		if err := ctx.ShouldBindBodyWith(&campaignCategory, binding.JSON); err == nil {
			campaigns := GetCampaignList(campaignCategory)
			ctx.JSON(http.StatusOK, campaigns)
		}
	}
}

func CampaignCreateHandler() func(*gin.Context) {
	return func(ctx *gin.Context) {
		campaign := Campaign{}

		if err := ctx.ShouldBindBodyWith(&campaign, binding.JSON); err == nil {
			campaign := GetCampaign(campaign)
			ctx.JSON(http.StatusOK, campaign)
		}
	}
}
