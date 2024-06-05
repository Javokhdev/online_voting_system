package service

import (
	"context"
	"log/slog"
	v "voting_service/genproto/voting"
	st "voting_service/storage/postgresql"
)

type CandidateService struct {
	storage st.Storage
	v.UnimplementedCandidateServiceServer
}

func NewCandidateService(storage *st.Storage) *CandidateService {
	return &CandidateService{
		storage: *storage,
	}
}

func (s *CandidateService) CreateCandidate(candidate *v.CreateCandidateReq) (*v.Void, error) {
	slog.Info("CreateCandidate Service called", "candidate", candidate)
	_, err := s.storage.CandidateS.CreateCandidate(candidate)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *CandidateService) GetCandidateById(ctx context.Context, id *v.ById) (*v.GetCandidateRes, error) {
	slog.Info("GetCandidateById Service called", "candidate_id", id)
	candidate, err := s.storage.CandidateS.GetCandidateById(id)

	if err != nil {
		return nil, err
	}

	return candidate, nil
}

func (s *CandidateService) GetAllCandidates(ctx context.Context,flt *v.Filter) (*v.GetAllCandidateRes, error) {
	slog.Info("GetAllCandidates Service called")

	res, err := s.storage.CandidateS.GetAllCandidates(flt)

	if err != nil{
		return nil, err
	}

	return res, nil
}

func (s *CandidateService) UpdateCandidate(ctx context.Context,candidate *v.GetCandidateRes) (*v.Void, error) {
	slog.Info("UpdateCandidate Service called", "candidate", candidate.GetId())
	_, err := s.storage.CandidateS.UpdateCandidate(candidate)

	return nil, err
}

func (s *CandidateService) DeleteCandidate(ctx context.Context,id *v.ById) (*v.Void, error) {
	slog.Info("DeleteCandidate Service called", "candidate_id", id.GetId())
	_, err := s.storage.CandidateS.DeleteCandidate(id)

	return nil, err
}
