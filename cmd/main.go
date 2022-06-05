package main

import (
	"github.com/gin-gonic/gin"
	"main/internal/config"
	"main/internal/router"
)

func main() {
	config.Init()
	defer config.GetConfig().Storage.Close()

	engine := gin.Default()
	router.Init(engine)
	engine.Run()
}
