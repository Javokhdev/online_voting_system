package api

import (
	"gateway/api/handler"

	"github.com/gin-gonic/gin"
	_ "gateway/docs"

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
	r.DELETE("/election/delete", handler.DeleteElection)

	r.POST("/candidate/create", handler.CreateCandidate)
	r.GET("/candidate/getbyid", handler.GetByIdCandidate)
	r.GET("/candidate/getall", handler.GetAllCandidates)
	r.PUT("/candidate/update", handler.UpdateCandidate)
	r.DELETE("/candidate/delete", handler.DeleteCandidate)

	r.POST("/public_vote/create", handler.CreatePublicVote)
	r.GET("/public_vote/getbyid", handler.GetByIdPublicVote)
	r.GET("/public_vote/getall", handler.GetAllPublicVotes)
	r.GET("/vote/getbyid", handler.GetVote)
	r.GET("/vote/getall", handler.GetVotes)
	
	r.POST("/party/create", handler.CreateParty)
	r.GET("/party/getbyid", handler.GetByIdParty)
	r.GET("/party/getall", handler.GetAllParties)
	r.PUT("/party/update", handler.UpdateParty)
	r.DELETE("/party/delete", handler.DeleteParty)

	r.POST("/public/create", handler.CreatePublic)
	r.GET("/public/getbyid", handler.GetByIdPublic)
	r.GET("/public/getall", handler.GetAllPublics)
	r.PUT("/public/update", handler.UpdatePublic)
	r.DELETE("/public/delete", handler.DeletePublic)

	
	// r.DELETE("/public_vote/delete", handler.DeleteCandidate)
	// r.POST("/candidate/create", handler.CreateCandidate)
	return r
}
