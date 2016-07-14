package main

import (
	"fmt"
	"os"

	"github.com/h-yamada/fb-bot-sample/config"
	"github.com/h-yamada/fb-bot-sample/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	config.FbToken = os.Getenv("FBTOKEN")
	config.MovikumaEsUrl = os.Getenv("ESURL")
	if config.FbToken != "" {
		router := gin.Default()

		router.GET("/webhook", handler.GetWebHook)
		router.POST("/webhook", handler.PostWebHook)

		router.Run(":9000")
	} else {
		fmt.Println("need to export FBTOKEN ")
	}
}
