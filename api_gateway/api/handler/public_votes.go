package handler

import (
	"gateway/genproto/public"
	"gateway/genproto/voting"
	"log/slog"
	"net/http"

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
	}
	id := public.ById{}
	id.Id = pv.GetPublicId()
	public, err := h.srvs.PublicService.GetById(ctx, &id)

	if err != nil || public.GetId() != id.GetId(){
		slog.Info("error with public id.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
		return
	}

	_, err = h.srvs.PublicVoteService.Create(ctx, &pv)

	if err != nil {
		slog.Info("Error publiuc_vote binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return
	}

	ctx.JSON(200, pv)
}

// @Summary Get Public_votes
// @Description Endpoint for getting public_votes
// @Tags Public_vote
// @Accept json
// @Produce json
// @Success 200 {object} voting.GetPublicVoteRes "Successfully getted public_votes"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get public_votes"
// @Router /public_vote/getall [GET]
func (h *Handler) GetAllPublicVotes(ctx *gin.Context) {
	filter := voting.Filter{}

	pvs, err := h.srvs.PublicVoteService.GetAll(ctx, &filter)

	if err != nil {
		slog.Info("error public_vote geting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return
	}

	ctx.JSON(200, pvs)
}

// @Summary Delete Public_vote
// @Description Endpoint for deleting public_vote
// @Tags Public_vote
// @Accept json
// @Produce json
// @Param  id query voting.ById true "ID"
// @Success 200 {object} voting.Void "Successfully deleted public_vote"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete public_vote"
// @Router /public_vote/delete [DELETE]
func (h *Handler) DeletePublicVote(ctx *gin.Context) {
	id := voting.ById{}

	id.Id = ctx.Query("id")	

	_, err := h.srvs.PublicVoteService.Delete(ctx, &id)	

	if err != nil {
		slog.Info("error public_vote binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return
	}

	ctx.JSON(200, "Public vote successfully deleted")
}


