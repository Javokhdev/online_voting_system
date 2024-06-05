package service

import (
	"context"
	"log/slog"
	v "voting_service/genproto/voting"
	st "voting_service/storage/postgresql"
)

type ElectionService struct {
	storage st.Storage
	v.UnimplementedElectionServiceServer
}

func NewElectionService(storage *st.Storage) *ElectionService {
	return &ElectionService{
		storage: *storage,
	}
}

func (s *ElectionService) CreateElection(ctx context.Context, election *v.CreateElectionReq) (*v.Void, error) {
	slog.Info("CreateElection Service called", "election", election.GetName())
	_, err := s.storage.ElectionS.CreateElection(election)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ElectionService) GetElectionById(ctx context.Context, id *v.ById) (*v.GetElectionRes, error) {
	slog.Info("GetElectionById Service called", "election_id", id.GetId())
	election, err := s.storage.ElectionS.GetElectionById(id)

	if err != nil {
		return nil, err
	}

	return election, nil
}

func (s *ElectionService) GetAllElections(ctx context.Context, flt *v.Filter) (*v.GetAllElectionRes, error) {
	slog.Info("GetAllElections Service called")

	res, err := s.storage.ElectionS.GetAllElections(flt)

	if err != nil{
		return nil, err
	}

	return res, nil
}

func (s *ElectionService) UpdateElection(ctx context.Context,election *v.GetElectionRes) (*v.Void, error) {
	slog.Info("UpdateElection Service called", "election", election.GetId())
	_, err := s.storage.ElectionS.UpdateElection(election)

	return nil, err
}

func (s *ElectionService) DeleteElection(ctx context.Context, id *v.ById) (*v.Void, error) {
	slog.Info("DeleteElection Service called", "election_id", id)
	_, err := s.storage.ElectionS.DeleteElection(id)

	return nil, err
}