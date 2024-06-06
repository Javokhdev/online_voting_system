package api

import (
	"gateway/api/handler"

	"github.com/gin-gonic/gin"
)

func NewGin(handler *handler.Handler) *gin.Engine {

	r := gin.Default()

	// r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.Use(Middleware())
	r.POST("/election/create", handler.CreateElection)
	// r.POST("/user/getall", handler.GetAllUser)
	// r.GET("/user/:id", handler.GetByIdUser)
	// r.POST("/courier", handler.Delivering)

	return r
}
