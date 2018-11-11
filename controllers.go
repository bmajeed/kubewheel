package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func listApps(context *gin.Context) {
	appsList, err := GetApps()
	if err != nil {
		panic(err.Error())
	}
	context.HTML(http.StatusOK, "index", gin.H{
		"apps": appsList.Items,
	})
}

func showApp(context *gin.Context) {
	deploymentsList, err := GetApp(context.Param("app_name"))
	if err != nil {
		panic(err.Error())
	}
	context.HTML(http.StatusOK, "deployment", gin.H{
		"deployments": deploymentsList.Items,
	})
}

func Register(group gin.RouterGroup) {
	group.GET("/", listApps)
	group.GET("/:app_name", showApp)
}
