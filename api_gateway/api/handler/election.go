package handler

import (
	"gateway/genproto/voting"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateElection(ctx *gin.Context) {
	es := voting.CreateElectionReq{}

	err := ctx.ShouldBindJSON(&es)

	if err != nil {
		slog.Info("error election binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
		return
	}

	_, err = h.srvs.ElectionService.Create(ctx, &es)

	if err != nil {
		slog.Info("error election binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return 
	}

	ctx.JSON(200, "Election successfully created")
}

