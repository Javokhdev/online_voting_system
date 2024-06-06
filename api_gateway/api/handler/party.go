package handler

import (
	"gateway/genproto/public"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)


// @Summary Create a new party
// @Description Endpoint for creating a new party
// @Name create_party
// @Slogan create_party
// @Opened_Date create_party
// @Description create_party
// @Tags Party
// @Accept json
// @Produce json
// @Param party body public.CreatePartyReq true "Party creation request payload"
// @Success 200 {object} public.Void "Successfully created party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create party"
// @Router /party/create [POST]
func (h *Handler) CreateParty(ctx *gin.Context) {
	py := public.CreatePartyReq{}

	err := ctx.ShouldBindJSON(&py)

	if err != nil {
		slog.Info("error party binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
		return
	}

	_, err = h.srvs.PartyService.Create(ctx, &py)

	if err != nil {
		slog.Info("error party binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return
	}

	ctx.JSON(200, "Party successfully created")
}
