package routes

import (
	"github.com/gin-gonic/gin"

	"hukum-onlen-go/internal/http/handlers"
)

// NewPublicRoutes registers public API routes.
func NewPublicRoutes(engine *gin.Engine) {
	pingHandler := handlers.NewPingHandler()

	engine.GET("/ping", pingHandler.Ping)
}
