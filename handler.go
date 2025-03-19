package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) processSentimentRequest(c *gin.Context) {
	var req SentimentRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	result, err := h.service.AnalyzeSentiment(req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to analyze sentiment",
			"details": err.Error(),
		})
		return
	}

	// No need to manually marshal, gin.JSON handles it
	c.JSON(http.StatusOK, result)
}

// AnalyzeSentiment handles sentiment analysis requests
func (h *Handler) AnalyzeSentiment(c *gin.Context) {
	h.processSentimentRequest(c)
}

// CallModelEndpoint is an alias for AnalyzeSentiment
// TODO: Consider deprecating this endpoint to avoid duplication
func (h *Handler) CallModelEndpoint(c *gin.Context) {
	h.processSentimentRequest(c)
}

// HealthCheck returns the API health status
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"version": "1.0", // Added version info
	})
}
