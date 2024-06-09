package handler

import (
	"gateway/genproto/public"
	"gateway/genproto/voting"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new public_vote
// @Description Endpoint for creating a new public_vote
// @Election_id create_public_vote
// @Public_id create_public_vote
// @Party_id create_public_vote
// @Tags Public_vote
// @Accept json
// @Produce json
// @Param public_vote body voting.CreatePublicVoteReq true "Public_vote creation request payload"
// @Success 200 {object} voting.Void "Successfully created public_vote"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create public_vote"
// @Router /public_vote/create [POST]
func (h *Handler) CreatePublicVote(ctx *gin.Context) {
	pv := voting.CreatePublicVoteReq{}

	err := ctx.ShouldBindJSON(&pv)

	if err != nil {
		slog.Info("Error publiuc_vote binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	id := public.ById{}
	id.Id = pv.GetPublicId()
	public, err := h.srvs.PublicService.GetById(ctx, &id)

	if err != nil || public.GetId() != id.GetId() {
		slog.Info("error with public id.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.srvs.PublicVoteService.Create(ctx, &pv)

	if err != nil {
		slog.Info("Error publiuc_vote binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, "Public vote successfully created")
}

// @Summary Get Public_vote
// @Description Endpoint for getting public_vote
// @Tags Public_vote
// @Accept json
// @Produce json
// @Param  id query voting.ById true "ID"
// @Success 200 {object} voting.GetPublicVoteRes "Successfully getted public_vote"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get public_vote"
// @Router /public_vote/getbyid [GET]
func (h *Handler) GetByIdPublicVote(ctx *gin.Context) {
	id := voting.ById{}

	id.Id = ctx.Query("id")

	pv, err := h.srvs.PublicVoteService.GetById(ctx, &id)

	if err != nil {
		slog.Info("error public_vote geting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, pv)
}

// @Summary Get Public_votes
// @Description Endpoint for getting public_votes
// @Tags Public_vote
// @Accept json
// @Produce json
// @Param public_vote query voting.Filter true "Public_vote getting request payload"
// @Success 200 {object} voting.GetPublicVoteRes "Successfully getted public_votes"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get public_votes"
// @Router /public_vote/getall [GET]
func (h *Handler) GetAllPublicVotes(ctx *gin.Context) {
	filter := voting.Filter{}

	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))

	filter.Limit = int32(limit)
	filter.Offset = int32(offset)

	pvs, err := h.srvs.PublicVoteService.GetAll(ctx, &filter)

	if err != nil {
		slog.Info("error public_vote geting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, pvs)
}

// @Summary Get Vote
// @Description Endpoint for getting vote
// @Tags Vote
// @Accept json
// @Produce json
// @Param vote query voting.ById true "Vote getting request payload"
// @Success 200 {object} voting.GetVoteById "Successfully getted vote"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get vote"
// @Router /vote/getbyid [GET]
func (h *Handler) GetVote(ctx *gin.Context) {
	id := voting.ById{}

	id.Id = ctx.Query("id")

	pv, err := h.srvs.PublicVoteService.GetVote(ctx, &id)

	if err != nil {
		slog.Info("error vote geting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, pv)
}

// @Summary Get Votes
// @Description Endpoint for getting vote
// @Tags Vote
// @Accept json
// @Produce json
// @Param vote query voting.Filter true "Votes getting request payload"
// @Success 200 {object} voting.GetAllVotes "Successfully getted vote"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get vote"
// @Router /vote/getall [GET]
func (h *Handler) GetVotes(ctx *gin.Context) {

	filter := voting.Filter{}

	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))

	filter.Limit = int32(limit)
	filter.Offset = int32(offset)

	vt, err := h.srvs.PublicVoteService.GetVotes(ctx, &filter)

	if err != nil {
		slog.Info("error vote geting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, vt)
}

// // @Summary Delete Public_vote
// // @Description Endpoint for deleting public_vote
// // @Tags Public_vote
// // @Accept json
// // @Produce json
// // @Param  id query voting.ById true "ID"
// // @Success 200 {object} voting.Void "Successfully deleted public_vote"
// // @Failure 400 {object} string "Invalid request payload"
// // @Failure 500 {object} string "Failed to delete public_vote"
// // @Router /public_vote/delete [DELETE]
// func (h *Handler) DeletePublicVote(ctx *gin.Context) {
// 	id := voting.ById{}

// 	id.Id = ctx.Query("id")

// 	_, err := h.srvs.PublicVoteService.Delete(ctx, &id)

// 	if err != nil {
// 		slog.Info("error public_vote binding.", "err", err.Error())
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(200, "Public vote successfully deleted")
// }
