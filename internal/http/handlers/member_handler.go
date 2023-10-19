package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"hukum-onlen-go/internal/usecases"
)

// CreateMemberBody represents the request body for creating member.
type CreateMemberBody struct {
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

// MemberHandler handles member-related HTTP requests.
type MemberHandler struct {
	useCase usecases.IMemberUsecase
}

// NewMemberHandler creates a new MemberHandler.
func NewMemberHandler(useCase usecases.IMemberUsecase) *MemberHandler {
	return &MemberHandler{
		useCase: useCase,
	}
}

// FindAllMembers finds all members.
func (h *MemberHandler) FindAllMembers(ctx *gin.Context) {
	emailFilter := ctx.Query("email")

	res, err := h.useCase.FindAllMembers(ctx, emailFilter)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to retrieve members"})
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(res, nil))
}

// FindMemberByID finds a member by ID.
func (h *MemberHandler) FindMemberByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := h.useCase.FindMemberByID(ctx, id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to retrieve member"})
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse(res, nil))
}

// CreateMember creates a new member.
func (h *MemberHandler) CreateMember(ctx *gin.Context) {
	body := &CreateMemberBody{}
	if err := ctx.ShouldBindJSON(body); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.useCase.CreateMember(ctx, body.Email, body.FirstName, body.LastName); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create member"})
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse(true, nil))
}
