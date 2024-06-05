package service

import (
	vote "voting_service/genproto/voting"
	st "voting_service/storage/postgresql"
	"log/slog"
)

type VoteService struct {
	storage *st.VoteStorage
	vote.UnimplementedPublicVoteServiceServer
}

func NewVoteService(storage *st.VoteStorage) *VoteService {
	return &VoteService{
		storage: storage,
	}
}

func (s *VoteService) CreateVote(Vote *vote.CreateVoteReq) (*vote.Void, error) {
	slog.Info("CreateVote Service called", "_vote", Vote)
	err := s.storage.CreateVote(Vote)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *VoteService) GetVoteById(id string) (*vote.GetVoteRes, error) {
	slog.Info("GetVoteById Service called", "vote_id", id)
	return s.storage.GetVoteById(id)
}

func (s *VoteService) GetAllVotes() (*vote.GetAllVoteRes, error) {
	slog.Info("GetAllVotes Service called")
	return s.storage.GetAllVotes()
}

func (s *VoteService) UpdateVote(Vote *vote.GetVoteRes) (*vote.Void, error) {
	slog.Info("UpdateVote Service called", "vote", Vote)
	err := s.storage.UpdateVote(Vote)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *VoteService) DeleteVote(id string) (*vote.Void, error) {
	slog.Info("DeleteVote Service called", "vote_id", id)
	err := s.storage.DeleteVote(id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
