package service

import (
	"context"
	"log/slog"
	st "public_service/storage/postgresql"
	p "public_service/genproto/public"
)

type PublicService struct {
	storage st.Storage
	p.UnimplementedPublicServiceServer
}

func NewPublicService(storage *st.Storage) *PublicService {
	return &PublicService{
		storage: *storage,
	}
}

func (s *PublicService) Create(ctx context.Context, public *p.CreatePublicReq) (*p.Void, error) {
	slog.Info("CreatePublic Service called", "public", public.GetFirstName())

	_, err := s.storage.PublicS.Create(public)

	return nil, err
}

func (s *PublicService) GetById(ctx context.Context, id *p.ById) (*p.GetPublicRes, error) {
	slog.Info("GetPublicById Service called", "public_id", id.GetId())
	party, err := s.storage.PublicS.GetById(id)

	if err != nil {
		return nil, err
	}

	return party, nil
}

func (s *PublicService) GetAll(ctx context.Context, flt *p.Filter) (*p.GetAllPublicRes, error) {
	slog.Info("GetAllPublics Service called")

	res, err := s.storage.PublicS.GetAll(flt)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PublicService) Update(ctx context.Context, public *p.GetPublicRes) (*p.Void, error) {
	slog.Info("UpdatePublic Service called", "public", public.GetId())
	_, err := s.storage.PublicS.Update(public)

	return nil, err
}

func (s *PublicService) Delete(ctx context.Context, id *p.ById) (*p.Void, error) {
	slog.Info("DeletePublic Service called", "public_id", id.GetId())
	_, err := s.storage.PublicS.Delete(id)

	return nil, err
}
