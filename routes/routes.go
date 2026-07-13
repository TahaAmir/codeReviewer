package routes

import (
	"codeReviewer/config"
	"codeReviewer/handlers"
	"codeReviewer/services"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, cfg *config.Config, aiSvc *services.AIService) {
	h := handlers.NewHandler(aiSvc)
	r.GET("/health", h.Health)
	r.POST("/review", h.Review)
}
