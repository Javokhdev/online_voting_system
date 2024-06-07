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

// @Summary Get Party
// @Description Endpoint for getting party
// @Tags Party
// @Accept json
// @Produce json
// @Param  id query public.ById true "ID"
// @Success 200 {object} public.GetPartyRes "Successfully getted party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get party"
// @Router /party/getbyid [GET]
func (h *Handler) GetByIdParty(ctx *gin.Context) {
	id := public.ById{}

	id.Id = ctx.Query("id")

	party, err := h.srvs.PartyService.GetById(ctx, &id)

	if err != nil {
		slog.Info("error party geting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return
	}

	ctx.JSON(200, party)
}


// @Summary Get All Parties
// @Description Endpoint for getting all parties
// @Tags Party
// @Accept json
// @Produce json
// @Param Party query public.Filter true "Parties geting request payload"
// @Success 200 {object} public.GetAllPartyRes "Successfully getted parties"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get parties"
// @Router /party/getall [GET]
func (h *Handler) GetAllParties(ctx *gin.Context) {
	filter := public.Filter{}

	parties, err := h.srvs.PartyService.GetAll(ctx, &filter)

	if err != nil {
		slog.Info("error parties geting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return
	}	

	ctx.JSON(200, parties)
}

// @Summary Update Party
// @Description Endpoint for deleting Party
// @Id update_party
// @Name update_party
// @Date update_party
// @Tags Party
// @Accept json
// @Produce json
// @Param party body public.GetPartyRes true "Party updating request payload"
// @Success 200 {object} public.Void "Successfully updated party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to update party"
// @Router /party/update [PUT]
func (h *Handler) UpdateParty(ctx *gin.Context) {
	py := public.GetPartyRes{}

	err := ctx.ShouldBindJSON(&py)

	if err != nil {
		slog.Info("error party binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
		return
	}

	_, err = h.srvs.PartyService.Update(ctx, &py)

	if err != nil {
		slog.Info("error party binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return
	}

	ctx.JSON(200, "Party successfully updated")
}

// @Summary Delete party
// @Description Endpoint for deleting party
// @Id delete_party
// @Tags Party
// @Accept json
// @Produce json
// @Param  id query public.ById true "ID"
// @Success 200 {object} public.Void "Successfully deleted party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete party"
// @Router /party/delete [DELETE]
func (h *Handler) DeleteParty(ctx *gin.Context) {
	party := public.ById{}

	party.Id = ctx.Query("id")

	_, err := h.srvs.PartyService.Delete(ctx, &party)

	if err != nil {
		slog.Info("error party binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "error in server"})
		return 
	}

	ctx.JSON(200, "Party successfully deleted")
}