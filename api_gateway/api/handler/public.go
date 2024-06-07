package handler

import (
	"gateway/genproto/public"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new public
// @Description Endpoint for creating a new public
// @First_Name create_public
// @Last_Name create_public
// @Birthday create_public
// @Gender create_public
// @Nation create_public
// @public_Id create_public
// @Tags Public
// @Accept json
// @Produce json
// @Param public body public.CreatePublicReq true "Public creation request payload"
// @Success 200 {object} public.Void "Successfully created public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create public"
// @Router /public/create [POST]
func (h *Handler) CreatePublic(ctx *gin.Context) {
	pb := public.CreatePublicReq{}

	err := ctx.ShouldBindJSON(&pb)

	if err != nil {
		slog.Info("error public binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
		return
	}

	_, err = h.srvs.PublicService.Create(ctx, &pb)	

	if err != nil {
		slog.Info("error public binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return 
	}

	ctx.JSON(200, "Public successfully created")
}


// @Summary Get Public
// @Description Endpoint for getting public
// @Tags Public
// @Accept json
// @Produce json
// @Param  id query public.ById true "ID"
// @Success 200 {object} public.GetPublicRes "Successfully getted public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get public"
// @Router /public/getbyid [GET]
func (h *Handler) GetByIdPublic(ctx *gin.Context) {
	id := public.ById{}

	id.Id = ctx.Query("id")

	public, err := h.srvs.PublicService.GetById(ctx, &id)

	if err != nil {
		slog.Info("error public geting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return
	}

	ctx.JSON(200, public)
}


// @Summary Get All Publics
// @Description Endpoint for getting all publics
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {object} public.GetAllPublicRes "Successfully getted publics"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get publics"
// @Router /public/getall [GET]
func (h *Handler) GetAllPublics(ctx *gin.Context) {	
	filter := public.Filter{}

	publics, err := h.srvs.PublicService.GetAll(ctx, &filter)

	if err != nil {
		slog.Info("error publics geting.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return
	}

	ctx.JSON(200, publics)
}																																																																			


// @Summary Update public
// @Description Endpoint for deleting public
// @Id update_public
// @First_Name update_public
// @Last_Name update_public
// @Birthday update_public
// @Gender update_public
// @Nation update_public
// @public_Id update_public
// @Tags Public
// @Accept json
// @Produce json
// @Param public body public.GetPublicRes true "Public updating request payload"
// @Success 200 {object} public.Void "Successfully updated public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to update public"
// @Router /public/update [PUT]
func (h *Handler) UpdatePublic(ctx *gin.Context) {
	pb := public.GetPublicRes{}

	err := ctx.ShouldBindJSON(&pb)

	if err != nil {
		slog.Info("error public binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not valid JSON"})
		return
	}

	_, err = h.srvs.PublicService.Update(ctx, &pb)
	
	if err != nil {
		slog.Info("error public binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return
	}

	ctx.JSON(200, "Public successfully updated")

}


// @Summary Delete public
// @Description Endpoint for deleting public
// @Id delete_public
// @First_Name delete_public
// @Last_Name delete_public
// @Birthday delete_public
// @Gender delete_public
// @Nation delete_public
// @public_Id delete_public
// @Tags Public
// @Accept json
// @Produce json
// @Param public body public.ById true "Public deleting request payload"
// @Success 200 {object} public.Void "Successfully deleted public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete public"
// @Router /public/delete [DELETE]
func (h *Handler) DeletePublic(ctx *gin.Context) {
	pb := public.ById{}

	pb.Id = ctx.Query("id")

	_, err := h.srvs.PublicService.Delete(ctx, &pb)

	if err != nil {
		slog.Info("error public binding.", "err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error in server"})
		return
	}

	ctx.JSON(200, "Public successfully deleted")
}