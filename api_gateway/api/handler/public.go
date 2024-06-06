package handler

import (
	"gateway/genproto/public"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

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