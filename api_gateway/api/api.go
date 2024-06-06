package api

import (
	"gateway/api/handler"

	"github.com/gin-gonic/gin"
	_ "gateway/docs"
	// "./docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Online Voting System Swagger UI
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name role

func NewGin(handler *handler.Handler) *gin.Engine {

	r := gin.Default()

	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/election/create", handler.CreateElection)
	r.GET("/election/getbyid", handler.GetByIdElection)
	r.GET("/election/getall", handler.GetAllElection)
	r.PUT("/election/update", handler.UpdateElection)
	// r.Use(Middleware())

	return r
}
