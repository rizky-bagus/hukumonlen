package middlewares

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	"hukum-onlen-go/internal/config"
)

// CorsMiddleware is a middleware for CORS
type CorsMiddleware struct {
	cfg    *config.Config
	engine *gin.Engine
}

// NewCorsMiddleware creates a new Cors middleware
func NewCorsMiddleware(cfg *config.Config, engine *gin.Engine) *CorsMiddleware {
	return &CorsMiddleware{
		cfg:    cfg,
		engine: engine,
	}
}

// Setup sets up the middleware
func (m *CorsMiddleware) Setup() {
	isDebugEnabled := m.cfg.Env == "development"
	m.engine.Use(cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool { return true },
		AllowedHeaders:  []string{"*"},
		AllowedMethods:  []string{"GET", "POST", "PUT", "PATCH", "HEAD", "OPTIONS"},
		Debug:           isDebugEnabled,
	}))
}
