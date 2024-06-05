package main

import (
	"log"
	"net"
	cf "voting_service/config"
	v "voting_service/genproto/voting"
	"voting_service/service"
	"voting_service/storage/postgresql"

	"google.golang.org/grpc"
)

func main() {

	config := cf.Load()

	db, err := postgresql.NewPostgresStorage(config)

	if err != nil {
		log.Fatal(err)
	}

	liss, err := net.Listen("tcp", config.TCP_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	v.RegisterElectionServiceServer(s, service.NewElectionService(db))
	v.RegisterCandidateServiceServer(s, service.NewCandidateService(db))
	v.RegisterPublicVoteServiceServer(s, service.NewPublicVoteService(db))

	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}