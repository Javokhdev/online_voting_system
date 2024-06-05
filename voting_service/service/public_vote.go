package service

import (
	"context"
	"log/slog"
	v "voting_service/genproto/voting"
	st "voting_service/storage/postgresql"
)

type PublicVoteService struct {
	storage st.Storage
	v.UnimplementedPublicVoteServiceServer
}

func NewPublicVoteService(storage *st.Storage) *PublicVoteService {
	return &PublicVoteService{
		storage: *storage,
	}
}

func (s *PublicVoteService) CreatePublicVote(ctx context.Context, publicVote *v.CreatePublicVoteReq) (*v.Void, error) {
	slog.Info("CreatePublicVote Service called", "public_vote", publicVote.GetElectionId())
	_, err := s.storage.PublicVoteS.CreatePublicVote(publicVote)

	return nil, err
}

func (s *PublicVoteService) GetPublicVoteById(ctx context.Context, id *v.ById) (*v.GetPublicVoteRes, error) {
	slog.Info("GetPublicVoteById Service called", "public_vote_id", id.GetId())
	
	res, err := s.storage.PublicVoteS.GetPublicVoteById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PublicVoteService) GetAllPublicVotes(ctx context.Context, flt *v.Filter) (*v.GetAllPublicVoteRes, error) {
	slog.Info("GetAllPublicVotes Service called")

	res, err := s.storage.PublicVoteS.GetAllPublicVotes(flt)

	if err != nil{
		return nil, err
	}

	return res, nil
}

func (s *PublicVoteService) UpdatePublicVote(ctx context.Context, publicVote *v.GetPublicVoteRes) (*v.Void, error) {
	slog.Info("UpdatePublicVote Service called", "public_vote", publicVote.GetId())

	_, err := s.storage.PublicVoteS.UpdatePublicVote(publicVote)

	return nil, err
}

func (s *PublicVoteService) DeletePublicVote(ctx context.Context, id *v.ById) (*v.Void, error) {
	slog.Info("DeletePublicVote Service called", "public_vote_id", id.GetId())

	_, err := s.storage.PublicVoteS.DeletePublicVote(id)

	return nil, err
}
