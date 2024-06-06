package handler

import (
	"gateway/genproto/public"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new public
// @Description Endpoint for creating a new public
// @First_Name create_public
// @Last_Name create_public
// @Birthday create_public
// @Gender create_public
// @Nation create_public
// @Party_Id create_public
// @Tags Public
// @Accept json
// @Produce json
// @Param party body public.CreatePublicReq true "Public creation request payload"
// @Success 200 {object} public.Void "Successfully created public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create public"
// @Router /public/create [POST]
func (h *Handler) CreatePublic(ctx *gin.Context) {
	pb := public.CreatePublicReq{}

	err := ctx.ShouldBindJSON(&pb)

	if err != nil {
		slog.Info("error public binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
		return
	}

	_, err = h.srvs.PublicService.Create(ctx, &pb)	

	if err != nil {
		slog.Info("error public binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return 
	}

	ctx.JSON(200, "Public successfully created")
}