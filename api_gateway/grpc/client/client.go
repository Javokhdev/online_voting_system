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
	connVot, err := grpc.NewClient(cfg.VOTING_SERVICE_HOST+cfg.VOTING_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connPub, err := grpc.NewClient(cfg.PUBLIC_SERVICE_HOST+cfg.PUBLIC_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		ElectionService:   voting.NewElectionServiceClient(connVot),
		CandidateService:  voting.NewCandidateServiceClient(connVot),
		PublicVoteService: voting.NewPublicVoteServiceClient(connVot),
		PartyService:      public.NewPartyServiceClient(connPub),
		PublicService:     public.NewPublicServiceClient(connPub),
	}, nil
}
