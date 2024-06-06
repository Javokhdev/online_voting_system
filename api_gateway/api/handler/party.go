package handler

import (
	"gateway/genproto/public"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateParty(ctx *gin.Context){
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