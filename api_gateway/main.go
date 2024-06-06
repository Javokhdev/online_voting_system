package main

import (
	"gateway/api"
	"gateway/api/handler"
	"gateway/config"
	"gateway/grpc/client"
	"log"
)

func main() {
	cfg := config.Load()

	services, err := client.NewGrpcClients(&cfg)
	if err != nil {
		log.Fatalf("error while connecting clients. err: %s", err.Error())
	}

	engine := api.NewGin(handler.NewHandler(services))

	err = engine.Run(cfg.HTTP_PORT)
	if err != nil {
		log.Fatalf("error while running server. err: %s", err.Error())
	}
}
