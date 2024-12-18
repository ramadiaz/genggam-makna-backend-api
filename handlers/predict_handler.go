package handlers

import (
	"genggam-makna-api/dto"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) ImagePredict(c *gin.Context) {
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "image required"})
		return
	}
	defer file.Close()

	image_data, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	result, err := h.service.ImagePredict(image_data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Body: result, Message: "image predicted successfully"})
}

func (h *compHandlers) VideoPredict(c *gin.Context) {
	file, _, err := c.Request.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "video required"})
		return
	}
	defer file.Close()

	video_data, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	result, err := h.service.VideoPredict(video_data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Body: result, Message: "video predicted successfully"})
}
