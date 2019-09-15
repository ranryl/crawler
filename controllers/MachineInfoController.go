package controllers

import (
	"crawler/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MachineInfoController 控制器
type MachineInfoController struct {
}

// Router Page
func (ctrl *MachineInfoController) Router(router *gin.Engine) {
	router.GET("/machine/info", ctrl.info)
	router.POST("/machine/info", ctrl.info)
}
func (ctrl *MachineInfoController) info(ctx *gin.Context) {
	miService := services.MachineInfoService{}
	data, err := miService.FindAll(1000)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// Before 前置操作
func (ctrl *MachineInfoController) Before() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (ctrl *MachineInfoController) index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "page/index.html", "")
}
func (ctrl *MachineInfoController) redirect(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
