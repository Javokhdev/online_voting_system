package service

import (
	cand "voting_service/genproto/candidate"
	st "voting_service/storage/postgresql"
	"log/slog"
)

type CandidateService struct {
	storage *st.CandidateStorage
	cand.UnimplementedCandidateServiceServer
}

func NewCandidateService(storage *st.CandidateStorage) *CandidateService {
	return &CandidateService{
		storage: storage,
	}
}

func (s *CandidateService) CreateCandidate(candidate *cand.CreateCandidateReq) (*cand.Void, error) {
	slog.Info("CreateCandidate Service called", "candidate", candidate)
	err := s.storage.CreateCandidate(candidate)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *CandidateService) GetCandidateById(id string) (*cand.GetCandidateRes, error) {
	slog.Info("GetCandidateById Service called", "candidate_id", id)
	return s.storage.GetCandidateById(id)
}

func (s *CandidateService) GetAllCandidates() (*cand.GetAllCandidateRes, error) {
	slog.Info("GetAllCandidates Service called")
	return s.storage.GetAllCandidates()
}

func (s *CandidateService) UpdateCandidate(candidate *cand.GetCandidateRes) (*cand.Void, error) {
	slog.Info("UpdateCandidate Service called", "candidate", candidate)
	err := s.storage.UpdateCandidate(candidate)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *CandidateService) DeleteCandidate(id string) (*cand.Void, error) {
	slog.Info("DeleteCandidate Service called", "candidate_id", id)
	err := s.storage.DeleteCandidate(id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
