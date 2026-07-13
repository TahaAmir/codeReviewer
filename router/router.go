package router

import (
	"codeReviewer/handlers"
	"codeReviewer/services"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, aiSvc *services.AIService) {
	h := handlers.NewHandler(aiSvc)
	api := r.Group("/api")
	{
		api.GET("/health", h.Health)
		api.POST("/review", h.Review)
	}
}
