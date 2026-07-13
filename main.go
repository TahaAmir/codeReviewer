package main

import (
	"codeReviewer/config"
	"codeReviewer/router"
	"codeReviewer/routes"
	"codeReviewer/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	aiSvc := services.New(cfg.OpenAiKey)
	r := gin.Default()
	router.Setup(r, aiSvc)
	routes.Register(r, cfg, aiSvc)
	log.Println("Server running at 8080 ")
	r.Run(":8080")
}
