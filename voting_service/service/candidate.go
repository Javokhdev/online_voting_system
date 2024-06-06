package service

import (
	"context"
	"errors"
	"log/slog"
	v "voting_service/genproto/voting"
	st "voting_service/storage/postgresql"
	// st "voting_service/storage"
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

func (s *CandidateService) Create(ctx context.Context, candidate *v.CreateCandidateReq) (*v.Void, error) {
	slog.Info("CreateCandidate Service called", "candidate", candidate)

	check := s.storage.CandidateS.CheckParty(candidate)

	if !check {
		err := errors.New("error while creating")
		return nil, err
	}

	check = s.storage.CandidateS.CheckPublic(candidate)

	if !check {
		err := errors.New("error while creating")
		return nil, err
	}

	_, err := s.storage.CandidateS.Create(candidate)

	return nil, err
}

func (s *CandidateService) GetById(ctx context.Context, id *v.ById) (*v.GetCandidateRes, error) {
	slog.Info("GetCandidateById Service called", "candidate_id", id)
	candidate, err := s.storage.CandidateS.GetById(id)

	if err != nil {
		return nil, err
	}

	return candidate, nil
}

func (s *CandidateService) GetAll(ctx context.Context, flt *v.Filter) (*v.GetAllCandidateRes, error) {
	slog.Info("GetAllCandidates Service called")

	res, err := s.storage.CandidateS.GetAll(flt)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CandidateService) Update(ctx context.Context, candidate *v.GetCandidateRes) (*v.Void, error) {
	slog.Info("UpdateCandidate Service called", "candidate", candidate.GetId())
	_, err := s.storage.CandidateS.Update(candidate)

	return nil, err
}

func (s *CandidateService) Delete(ctx context.Context, id *v.ById) (*v.Void, error) {
	slog.Info("DeleteCandidate Service called", "candidate_id", id.GetId())
	_, err := s.storage.CandidateS.Delete(id)

	return nil, err
}
