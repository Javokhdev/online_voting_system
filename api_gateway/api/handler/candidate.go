package handler

import (
	"gateway/genproto/voting"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCandidate(ctx *gin.Context) {
	cs := voting.CreateCandidateReq{}

	err := ctx.ShouldBindJSON(&cs)

	if err != nil {
		slog.Info("error candidate binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
		return
	}

	_, err = h.srvs.CandidateService.Create(ctx, &cs)

	if err != nil {
		slog.Info("error candidate binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return 
	}

	ctx.JSON(200, "Candidate successfully created")
}