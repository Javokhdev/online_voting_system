package service

import (
	elec "voting_service/genproto/election"
	st "voting_service/storage/postgresql"
	"log/slog"
)

type ElectionService struct {
	storage *st.ElectionStorage
	elec.UnimplementedElectionServiceServer
}

func NewElectionService(storage *st.ElectionStorage) *ElectionService {
	return &ElectionService{
		storage: storage,
	}
}

func (s *ElectionService) CreateElection(election *elec.CreateElectionReq) (*elec.Void, error) {
	slog.Info("CreateElection Service called", "election", election)
	err := s.storage.CreateElection(election)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ElectionService) GetElectionById(id string) (*elec.GetElectionRes, error) {
	slog.Info("GetElectionById Service called", "election_id", id)
	return s.storage.GetElectionById(id)
}

func (s *ElectionService) GetAllElections() (*elec.GetAllElectionRes, error) {
	slog.Info("GetAllElections Service called")
	return []s.storage.GetAllElections()
}

func (s *ElectionService) UpdateElection(election *elec.GetElectionRes) (*elec.Void, error) {
	slog.Info("UpdateElection Service called", "election", election)
	err := s.storage.UpdateElection(election)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *ElectionService) DeleteElection(id string) (*elec.Void, error) {
	slog.Info("DeleteElection Service called", "election_id", id)
	err := s.storage.DeleteElection(id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

