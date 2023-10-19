package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	entity "hukum-onlen-go/entity"
	"hukum-onlen-go/internal/usecases"
)

// CreateInvitationStatusBody represents the request body for creating invitation status.
type CreateInvitationStatusBody struct {
	Name string `json:"name" binding:"required"`
}

// CreateGatheringTypeBody represents the request body for creating gathering type.
type CreateGatheringTypeBody struct {
	Name string `json:"name" binding:"required"`
}

// CreateGatheringBody represents the request body for creating gathering.
type CreateGatheringBody struct {
	Name        string        `json:"name" binding:"required"`
	Location    string        `json:"location" binding:"required"`
	CreatorID   entity.UUID   `json:"creator_id" binding:"required,uuid"`
	TypeID      entity.UUID   `json:"type_id" binding:"required,uuid"`
	ScheduledAt time.Time     `json:"scheduled_at" binding:"required"`
	AttendeeIDs []entity.UUID `json:"attendee_ids"`
}

// GatheringHandler handles gathering-related HTTP requests.
type GatheringHandler struct {
	useCase usecases.IGatheringUsecase
}

// NewGatheringHandler creates a new GatheringHandler.
func NewGatheringHandler(useCase usecases.IGatheringUsecase) *GatheringHandler {
	return &GatheringHandler{
		useCase: useCase,
	}
}

// FindAllInvitationStatuses finds all invitation statuses.
func (h *GatheringHandler) FindAllInvitationStatuses(ctx *gin.Context) {
	invitationStatuses, err := h.useCase.FindAllInvitationStatuses(ctx)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to retrieve invitation status"})
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(invitationStatuses, nil))
}

// FindInvitationStatusByID finds an invitation status by ID.
func (h *GatheringHandler) FindInvitationStatusByID(ctx *gin.Context) {
	id := ctx.Param("id")

	invitationStatus, err := h.useCase.FindInvitationStatusByID(ctx, id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to retrieve invitation status"})
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(invitationStatus, nil))
}

// CreateInvitationStatus creates a gathering type.
func (h *GatheringHandler) CreateInvitationStatus(ctx *gin.Context) {
	body := &CreateInvitationStatusBody{}
	if err := ctx.ShouldBindJSON(body); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.useCase.CreateInvitationStatus(ctx, body.Name); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create invitation status"})
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(true, nil))
}

// FindAllGatheringTypes finds all gathering types.
func (h *GatheringHandler) FindAllGatheringTypes(ctx *gin.Context) {
	res, err := h.useCase.FindAllGatheringTypes(ctx)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to retrieve gathering type"})
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(res, nil))
}

// FindGatheringTypeByID finds a gathering type by ID.
func (h *GatheringHandler) FindGatheringTypeByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := h.useCase.FindGatheringTypeByID(ctx, id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to retrieve gathering type"})
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(res, nil))
}

// CreateGatheringType creates a gathering type.
func (h *GatheringHandler) CreateGatheringType(ctx *gin.Context) {
	body := &CreateGatheringTypeBody{}
	if err := ctx.ShouldBindJSON(body); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.useCase.CreateGatheringType(ctx, body.Name); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create gathering type"})
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(true, nil))
}

// CreateGathering creates a gathering.
func (h *GatheringHandler) CreateGathering(ctx *gin.Context) {
	body := &CreateGatheringBody{}
	if err := ctx.ShouldBindJSON(body); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.useCase.CreateGathering(
		ctx,
		body.Name,
		body.Location,
		body.CreatorID,
		body.TypeID,
		body.ScheduledAt,
		body.AttendeeIDs,
	); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create gathering"})
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(true, nil))
}
