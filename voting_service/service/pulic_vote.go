package service

import (
	vote "voting_service/genproto/voting"
	st "voting_service/storage/postgresql"
	"log/slog"
)

type PublicVoteService struct {
	storage *st.PublicVoteStorage
	vote.UnimplementedPublicVoteServiceServer
}

func NewPublicVoteService(storage *st.PublicVoteStorage) *PublicVoteService {
	return &PublicVoteService{
		storage: storage,
	}
}

func (s *PublicVoteService) CreatePublicVote(publicVote *vote.CreatePublicVoteReq) (*vote.Void, error) {
	slog.Info("CreatePublicVote Service called", "public_vote", publicVote)
	err := s.storage.CreatePublicVote(publicVote)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *PublicVoteService) GetPublicVoteById(id string) (*vote.GetPublicVoteRes, error) {
	slog.Info("GetPublicVoteById Service called", "public_vote_id", id)
	return s.storage.GetPublicVoteById(id)
}

func (s *PublicVoteService) GetAllPublicVotes() (*vote.GetAllPublicVoteRes, error) {
	slog.Info("GetAllPublicVotes Service called")
	return s.storage.GetAllPublicVotes()
}

func (s *PublicVoteService) UpdatePublicVote(publicVote *vote.GetPublicVoteRes) (*vote.Void, error) {
	slog.Info("UpdatePublicVote Service called", "public_vote", publicVote)
	err := s.storage.UpdatePublicVote(publicVote)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *PublicVoteService) DeletePublicVote(id string) (*vote.Void, error) {
	slog.Info("DeletePublicVote Service called", "public_vote_id", id)
	err := s.storage.DeletePublicVote(id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}


