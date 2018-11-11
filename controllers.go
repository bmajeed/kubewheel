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
	deploymentsList, err := App{context.Param("app_name")}.GetAppDeployments()
	if err != nil {
		panic(err.Error())
	}
	context.HTML(http.StatusOK, "deployment", gin.H{
		"deployments": deploymentsList.Items,
	})
}

func createApp(context *gin.Context) {
	var app App
	err := context.Bind(&app)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	err = app.Create()
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	context.Redirect(http.StatusMovedPermanently, "/apps/"+app.Name)
}

func Register(group gin.RouterGroup) {
	group.GET("/", listApps)
	group.POST("/", createApp)
	group.GET("/:app_name", showApp)
}
