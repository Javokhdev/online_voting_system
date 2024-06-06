package service

import (
	"context"
	"log/slog"
	v "voting_service/genproto/voting"
	st "voting_service/storage/postgresql"
	// st "voting_service/storage"
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

func (s *PublicVoteService) Create(ctx context.Context, publicVote *v.CreatePublicVoteReq) (*v.Void, error) {
	slog.Info("CreatePublicVote Service called", "public_vote", publicVote.GetElectionId())
	_, err := s.storage.PublicVoteS.Create(publicVote)

	return nil, err
}

func (s *PublicVoteService) GetById(ctx context.Context, id *v.ById) (*v.GetPublicVoteRes, error) {
	slog.Info("GetPublicVoteById Service called", "public_vote_id", id.GetId())
	
	res, err := s.storage.PublicVoteS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PublicVoteService) GetAll(ctx context.Context, flt *v.Filter) (*v.GetAllPublicVoteRes, error) {
	slog.Info("GetAllPublicVotes Service called")

	res, err := s.storage.PublicVoteS.GetAll(flt)

	if err != nil{
		return nil, err
	}

	return res, nil
}

func (s *PublicVoteService) Update(ctx context.Context, publicVote *v.GetPublicVoteRes) (*v.Void, error) {
	slog.Info("UpdatePublicVote Service called", "public_vote", publicVote.GetId())

	_, err := s.storage.PublicVoteS.Update(publicVote)

	return nil, err
}

func (s *PublicVoteService) Delete(ctx context.Context, id *v.ById) (*v.Void, error) {
	slog.Info("DeletePublicVote Service called", "public_vote_id", id.GetId())

	_, err := s.storage.PublicVoteS.Delete(id)

	return nil, err
}

func (s *PublicVoteService) GetVote(ctx context.Context, vote *v.ById)(*v.GetVoteById, error){
	slog.Info("Get Vote by candidate id", "vote byid",vote.GetId())

	res, err := s.storage.PublicVoteS.GetVote(vote)

	if err != nil{
		return nil, err
	}

	return res, nil
}

func (s *PublicVoteService) GetVotes(ctx context.Context, flt *v.Filter)(*v.GetAllVotes, error){
	slog.Info("Get all votes with count")

	res, err := s.storage.PublicVoteS.GetVotes(flt)

	if err != nil{
		return nil, err
	}

	return res, nil
}