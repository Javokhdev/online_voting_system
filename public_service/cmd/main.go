package main

import (
	"log"
	"net"
	cf "public_service/config"
	p "public_service/genproto/public"
	"public_service/service"
	"public_service/storage/postgresql"

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

	p.RegisterPublicServiceServer(s, service.NewPublicService(db))
	p.RegisterPartyServiceServer(s, service.NewPartyService(db))

	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}