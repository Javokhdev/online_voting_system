package handler

import (
	"gateway/grpc/client"
)

type Handler struct {
	srvs *client.GrpcClients
}

func NewHandler(srvs *client.GrpcClients) *Handler {
	return &Handler{srvs: srvs}
}