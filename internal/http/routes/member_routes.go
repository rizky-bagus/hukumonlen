package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"

	"hukum-onlen-go/internal/http/handlers"
	"hukum-onlen-go/internal/repositories"
	"hukum-onlen-go/internal/usecases"
)

// NewMemberRoutes registers gathering-related routes.
func NewMemberRoutes(db *bun.DB, engine *gin.Engine) {
	memberRepository := repositories.NewMemberRepository(db)
	memberUsecase := usecases.NewMemberUsecase(memberRepository)
	memberHandler := handlers.NewMemberHandler(memberUsecase)

	member := engine.Group("/members")
	member.GET("/", memberHandler.FindAllMembers)
	member.GET("/:id", memberHandler.FindMemberByID)
	member.POST("/", memberHandler.CreateMember)
}
