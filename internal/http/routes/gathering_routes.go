package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"

	"hukum-onlen-go/internal/http/handlers"
	repositories "hukum-onlen-go/internal/repositories"
	"hukum-onlen-go/internal/usecases"
)

// NewGatheringRoutes registers gathering-related routes.
func NewGatheringRoutes(bunDB *bun.DB, engine *gin.Engine) {
	memberRepository := repositories.NewMemberRepository(bunDB)
	invitationStatusRepository := repositories.NewInvitationStatusRepository(bunDB)
	gatheringTypeRepository := repositories.NewGatheringTypeRepository(bunDB)
	invitationRepository := repositories.NewInvitationRepository(bunDB)
	gatheringRepository := repositories.NewGatheringRepository(bunDB)
	gatheringUsecase := usecases.NewGatheringUsecase(
		memberRepository, invitationStatusRepository, gatheringTypeRepository, invitationRepository, gatheringRepository)

	gatheringHandler := handlers.NewGatheringHandler(gatheringUsecase)

	invStatus := engine.Group("/invitation-statuses")
	invStatus.GET("/", gatheringHandler.FindAllInvitationStatuses)
	invStatus.GET("/:id", gatheringHandler.FindInvitationStatusByID)
	invStatus.POST("/", gatheringHandler.CreateInvitationStatus)

	gatheringType := engine.Group("/gathering-types")
	gatheringType.GET("/", gatheringHandler.FindAllGatheringTypes)
	gatheringType.GET("/:id", gatheringHandler.FindGatheringTypeByID)
	gatheringType.POST("/", gatheringHandler.CreateGatheringType)

	gathering := engine.Group("/gatherings")
	gathering.POST("/", gatheringHandler.CreateGathering)
}
