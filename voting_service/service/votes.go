package service

import (
	"context"
	"log/slog"
	v "voting_service/genproto/voting"
	st "voting_service/storage/postgresql"
)

type VoteService struct {
	storage st.Storage
	v.UnimplementedPublicVoteServiceServer
}

func NewVoteService(storage *st.Storage) *VoteService {
	return &VoteService{
		storage: *storage,
	}
}

func (s *VoteService) CreateVote(ctx context.Context, Vote *v.CreateVoteReq) (*v.Void, error) {
	slog.Info("CreateVote Service called", "_vote", Vote.GetCandidateId())
	err := s.storage.VotesS.CreateVote(Vote)

	return nil, err
}

func (s *VoteService) GetVoteById(ctx context.Context, id *v.ById) (*v.GetVotesRes, error) {
	slog.Info("GetVoteById Service called", "vote_id", id.GetId())
	
	res, err := s.storage.VotesS.GetVoteById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *VoteService) GetAllVotes(ctx context.Context, flt *v.Filter) (*v.GetAllVotesRes, error) {
	slog.Info("GetAllVotes Service called")
	
	res, err := s.storage.VotesS.GetAllVotes(flt)

	if err != nil{
		return nil, err
	}

	return res, nil
}

func (s *VoteService) UpdateVote(ctx context.Context, Vote *v.GetVotesRes) (*v.Void, error) {
	slog.Info("UpdateVote Service called", "vote", Vote.GetId())
	_, err := s.storage.VotesS.UpdateVote(Vote)

	return nil, err
}

func (s *VoteService) DeleteVote(ctx context.Context, id *v.ById) (*v.Void, error) {
	slog.Info("DeleteVote Service called", "vote_id", id.GetId())
	_, err := s.storage.VotesS.DeleteVote(id)

	return nil, err
}
