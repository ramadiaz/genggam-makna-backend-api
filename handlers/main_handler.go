package handlers

import (
	"net/http"
	"genggam-makna-api/dto"
	"genggam-makna-api/services"

	"github.com/gin-gonic/gin"
)

type compHandlers struct {
	service services.CompService
}

func NewCompHandlers(s services.CompService) *compHandlers {
	return &compHandlers{
		service: s,
	}
}

func (h *compHandlers) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "pong"})
}
