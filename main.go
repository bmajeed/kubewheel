package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kubewheel/GinHTMLRender"
	"net/http"
	"os"
)

var VERSION = "v1.0"

func main() {
	r := gin.Default()
	render := GinHTMLRender.New()
	render.TemplatesDir = "templates/"
	render.Debug = gin.IsDebugging()
	r.HTMLRender = render.Create()

	r.GET("/status", func(context *gin.Context) {
		clientset, err := getKubeClientset()
		if err != nil{
			panic(err.Error())
		}
		info, _ := clientset.ServerVersion()
		context.JSON(http.StatusOK, gin.H{
			"kubernetes": info,
			"KubeWheel": gin.H{
				"version": VERSION,
			},
		})
	})
	Register(*r.Group("/apps"))

	err := checkEnv()
	if err != nil {
		panic(err.Error())
	}
	r.Run()
}

func checkEnv() error {
	if os.Getenv("GITHUB_TOKEN") == ""{
		return errors.New("GITHUB_TOKEN token is not set")
	}
	clientset, err := getKubeClientset()
	if err != nil {
		return err
	}
	_, err = clientset.ServerVersion()
	if err != nil {
		return err
	}
	return nil
}
