package storage

import (
	"database/sql"
	v "voting_service/genproto/voting"
)

type Storage interface {
	Db() *sql.DB
	ElectionS() ElectionI
	CandidateS() CandidateI
	PublicVoteS() PublicVoteI
}

type ElectionI interface {
	Create(*v.CreateElectionReq) (*v.Void, error)
	GetById(*v.ById) (*v.GetElectionRes, error)
	GetAll(*v.Filter) (*v.GetAllElectionRes, error)
	Update(*v.GetElectionRes) (*v.Void, error)
	Delete(*v.ById) (*v.Void, error)
}

type CandidateI interface {
	Create(*v.CreateCandidateReq) (*v.Void, error)
	GetById(*v.ById) (*v.GetCandidateRes, error)
	GetAll(*v.Filter) (*v.GetAllCandidateRes, error)
	Update(*v.GetCandidateRes) (*v.Void, error)
	Delete(*v.ById) (*v.Void, error)
}

type PublicVoteI interface {
	Create(*v.CreatePublicVoteReq) (*v.Void, error)
	GetById(*v.ById) (*v.GetPublicVoteRes, error)
	GetAll(*v.Filter) (*v.GetAllPublicVoteRes, error)
	Update(*v.GetPublicVoteRes) (*v.Void, error)
	Delete(*v.ById) (*v.Void, error)
	GetVote(*v.ById) (*v.GetVoteById, error)
	GetVotes(*v.Filter) (*v.GetAllVotes, error)
}
