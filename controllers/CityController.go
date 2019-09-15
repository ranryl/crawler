package controllers

import (
	"crawler/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CityController struct
type CityController struct{}

func (ctrl *CityController) getAll(ctx *gin.Context) {
	cityService := services.CityService{}
	cityData, err := cityService.FindAll(1000)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": cityData,
	})
}

// Router city router
func (ctrl *CityController) Router(router *gin.Engine) {
	router.GET("/city/getall", ctrl.getAll)
	router.POST("/city/getall", ctrl.getAll)
}
