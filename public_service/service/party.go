package service

import (
	"context"
	"log/slog"
	p "public_service/genproto/public"
	st "public_service/storage/postgresql"
)

type PartyService struct{
	storage st.Storage
	p.UnimplementedPartyServiceServer
}

func NewPartyService(storage *st.Storage) *PartyService{
	return &PartyService{
		storage: *storage,
	}
}

func (s *PartyService) Create(ctx context.Context, party *p.CreatePartyReq)(*p.Void, error){
	slog.Info("CreateParty Service called", "party", party.GetName())

	_, err := s.storage.PartyS.Create(party)

	return nil, err
}

func (s *PartyService) GetById(ctx context.Context, id *p.ById)(*p.GetPartyRes, error){
	slog.Info("GetPartyById Service called", "party_id", id.GetId())
	party, err := s.storage.PartyS.GetById(id)

	if err != nil {
		return nil, err
	}

	return party, nil
}

func (s *PartyService) GetAll(ctx context.Context, flt *p.Filter)(*p.GetAllPartyRes, error){
	slog.Info("GetAllParties Service called")

	res, err := s.storage.PartyS.GetAll(flt)

	if err != nil{
		return nil, err
	}

	return res, nil
}

func (s *PartyService) Update(ctx context.Context, party *p.GetPartyRes)(*p.Void, error){
	slog.Info("UpdateParty Service called", "party", party.GetId())
	_, err := s.storage.PartyS.Update(party)

	return nil, err
}

func (s *PartyService) Delete(ctx context.Context, id *p.ById)(*p.Void, error){
	slog.Info("DeleteParty Service called", "party_id", id.GetId())
	_, err := s.storage.PartyS.Delete(id)

	return nil, err
}