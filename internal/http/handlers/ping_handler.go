package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler handles HTTP requests for ping requests.
type PingHandler struct{}

// NewPingHandler creates a new PingHandler.
func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// Ping handles ping requests.
func (h *PingHandler) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
