package handlers

import (
	"codeReviewer/models"
	"codeReviewer/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	aiService *services.AIService
}

func NewHandler(aiService *services.AIService) *Handler {
	return &Handler{aiService: aiService}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "200",
	})
}

func (h *Handler) Review(c *gin.Context) {
	var req models.ReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"Status": "400",
		})
		return
	}

	res, err := h.aiService.AnalyzeCode(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
