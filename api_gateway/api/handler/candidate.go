package handler

import (
	"gateway/genproto/voting"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new candidate
// @Description Endpoint for creating a new candidate
// @Election_id create_candidate
// @Public_id create_candidate
// @Party_id create_candidate
// @Tags Candidate
// @Accept json
// @Produce json
// @Param candidate body voting.CreateCandidateReq true "Candidate creation request payload"
// @Success 200 {object} voting.Void "Successfully created candidate"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create candidate"
// @Router /candidate/create [POST]
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