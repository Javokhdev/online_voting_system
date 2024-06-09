package handler

import (
	"gateway/genproto/voting"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new election
// @Description Endpoint for creating a new election
// @Name create_election
// @Date create_election
// @Tags Election
// @Accept json
// @Produce json
// @Param election body voting.CreateElectionReq true "Election creation request payload"
// @Success 200 {object} voting.Void "Successfully created election"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create election"
// @Router /election/create [POST]
func (h *Handler) CreateElection(ctx *gin.Context) {
	es := voting.CreateElectionReq{}

	err := ctx.ShouldBindJSON(&es)

	if err != nil {
		slog.Info("error election binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.srvs.ElectionService.Create(ctx, &es)

	if err != nil {
		slog.Info("error election binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, "Election successfully created")
}

// @Summary Get Election
// @Description Endpoint for getting election
// @Tags Election
// @Accept json
// @Produce json
// @Param  id query voting.ById true "ID"
// @Success 200 {object} voting.GetElectionRes "Successfully getted election"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get election"
// @Router /election/getbyid [GET]
func (h *Handler) GetByIdElection(ctx *gin.Context) {
	id := voting.ById{}

	id.Id = ctx.Query("id")

	election, err := h.srvs.ElectionService.GetById(ctx, &id)

	if err != nil {
		slog.Info("error election geting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, election)
}

// @Summary Get All Elections
// @Description Endpoint for getting all elections
// @Tags Election
// @Accept json
// @Produce json
// @Param election query voting.Filter true "Elections geting request payload"
// @Success 200 {object} voting.GetAllElectionRes "Successfully getted elections"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get elections"
// @Router /election/getall [GET]
func (h *Handler) GetAllElection(ctx *gin.Context) {
	filter := voting.Filter{}

	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))

	filter.Limit = int32(limit)
	filter.Offset = int32(offset)

	elections, err := h.srvs.ElectionService.GetAll(ctx, &filter)

	if err != nil {
		slog.Info("error elections getting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, elections)
}

// @Summary Update election
// @Description Endpoint for deleting election
// @Id update_election
// @Name update_election
// @Date update_election
// @Tags Election
// @Accept json
// @Produce json
// @Param election body voting.GetElectionRes true "Election updaing request payload"
// @Success 200 {object} voting.Void "Successfully updated election"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to update election"
// @Router /election/update [PUT]
func (h *Handler) UpdateElection(ctx *gin.Context) {
	election := voting.GetElectionRes{}

	err := ctx.ShouldBindJSON(&election)

	if err != nil {
		slog.Info("error election binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.srvs.ElectionService.Update(ctx, &election)

	if err != nil {
		slog.Info("error election binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, "Election successfully updated")
}

// @Summary Delete election
// @Description Endpoint for deleting election
// @Id delete_election
// @Tags Election
// @Accept json
// @Produce json
// @Param  id query voting.ById true "ID"
// @Success 200 {object} voting.Void "Successfully deleted election"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete election"
// @Router /election/delete [DELETE]
func (h *Handler) DeleteElection(ctx *gin.Context) {
	election := voting.ById{}

	election.Id = ctx.Query("id")

	_, err := h.srvs.ElectionService.Delete(ctx, &election)

	if err != nil {
		slog.Info("error election binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, "Election deleted successfully")
}
