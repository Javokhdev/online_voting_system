package handler

import (
	"gateway/genproto/public"
	"gateway/genproto/voting"
	"log/slog"
	"net/http"
	"strconv"

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := public.ById{}
	id.Id = cs.GetPartyId()
	party, err := h.srvs.PartyService.GetById(ctx, &id)

	if err != nil || party.GetId() != id.GetId() {
		slog.Info("error with party id.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id.Id = cs.GetPublicId()
	public, err := h.srvs.PublicService.GetById(ctx, &id)

	if err != nil || public.GetId() != id.GetId() {
		slog.Info("error with public id.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.srvs.CandidateService.Create(ctx, &cs)

	if err != nil {
		slog.Info("error candidate binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, "Candidate successfully created")
}

// @Summary Get Candidate
// @Description Endpoint for getting candidate
// @Tags Candidate
// @Accept json
// @Produce json
// @Param  id query voting.ById true "ID"
// @Success 200 {object} voting.GetCandidateRes "Successfully getted candidate"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get candidate"
// @Router /candidate/getbyid [GET]
func (h *Handler) GetByIdCandidate(ctx *gin.Context) {
	id := voting.ById{}

	id.Id = ctx.Query("id")

	candidate, err := h.srvs.CandidateService.GetById(ctx, &id)

	if err != nil {
		slog.Info("error candidate geting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, candidate)
}

// @Summary Get All Candidates
// @Description Endpoint for getting all candidates
// @Tags Candidate
// @Accept json
// @Produce json
// @Param election query voting.Filter true "Candidates geting request payload"
// @Success 200 {object} voting.GetCandidateRes "Successfully getted candidates"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get candidates"
// @Router /candidate/getall [GET]
func (h *Handler) GetAllCandidates(ctx *gin.Context) {
	filter := voting.Filter{}

	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))

	filter.Limit = int32(limit)
	filter.Offset = int32(offset)

	candidates, err := h.srvs.CandidateService.GetAll(ctx, &filter)

	if err != nil {
		slog.Info("error candidate geting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, candidates)
}

// @Summary Update Candidate
// @Description Endpoint for deleting candidate
// @Id update_candidate
// @Election_id delete_candidate
// @Public_id delete_candidate
// @Party_id delete_candidate
// @Tags Candidate
// @Accept json
// @Produce json
// @Param candidate body voting.GetCandidateRes true "Candidate updaing request payload"
// @Success 200 {object} voting.Void "Successfully updated candidate"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to update candidate"
// @Router /candidate/update [PUT]
func (h *Handler) UpdateCandidate(ctx *gin.Context) {
	candidate := voting.GetCandidateRes{}

	err := ctx.ShouldBindJSON(&candidate)

	if err != nil {
		slog.Info("error candidate binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.srvs.CandidateService.Update(ctx, &candidate)

	if err != nil {
		slog.Info("error candidate binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, "Candidate successfully updated")
}

// @Summary Delete Candidate
// @Description Endpoint for deleting candidate
// @Id delete_candidate
// @Election_id delete_candidate
// @Public_id delete_candidate
// @Party_id delete_candidate
// @Tags Candidate
// @Accept json
// @Produce json
// @Param  id query voting.ById true "ID"
// @Success 200 {object} voting.Void "Successfully deleted candidate"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete candidate"
// @Router /candidate/delete [DELETE]
func (h *Handler) DeleteCandidate(ctx *gin.Context) {
	id := voting.ById{}

	id.Id = ctx.Query("id")

	_, err := h.srvs.CandidateService.Delete(ctx, &id)

	if err != nil {
		slog.Info("error candidate binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, "Candidate successfully deleted")

}
