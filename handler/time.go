package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getTime(c *gin.Context) {
	timeNow := time.Now()
	formatTime := timeNow.Format(time.RFC3339)
	c.JSON(http.StatusOK, gin.H{
		"time": formatTime,
	})
}

func (h *Handler) greet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from server",
	})
}
