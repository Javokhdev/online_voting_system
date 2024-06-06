package handler

import (
	"gateway/genproto/voting"
	"log/slog"
	"net/http"	

	"github.com/gin-gonic/gin"
)


func (h *Handler) CreatePublicVote(ctx *gin.Context) {
	pv := voting.CreatePublicVoteReq{}

	err := ctx.ShouldBindJSON(&pv)

	if err != nil {
		slog.Info("Error publiuc_vote binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
	}

	_, err = h.srvs.PublicVoteService.Create(ctx, &pv)

	if err != nil {
		slog.Info("Error publiuc_vote binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return 
	}

	ctx.JSON(200, "Public vote successfully created")
}
