package client

import (
	"gateway/config"
	"gateway/genproto/public"
	"gateway/genproto/voting"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	ElectionService   voting.ElectionServiceClient
	CandidateService  voting.CandidateServiceClient
	PublicVoteService voting.PublicVoteServiceClient
	PartyService      public.PartyServiceClient
	PublicService     public.PublicServiceClient
}

func NewGrpcClients(cfg *config.Config) (*GrpcClients, error) {
	conn, err := grpc.NewClient(cfg.SERVICE_HOST+cfg.SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		ElectionService:   voting.NewElectionServiceClient(conn),
		CandidateService:  voting.NewCandidateServiceClient(conn),
		PublicVoteService: voting.NewPublicVoteServiceClient(conn),
		PartyService:      public.NewPartyServiceClient(conn),
		PublicService:     public.NewPublicServiceClient(conn),
	}, nil
}
