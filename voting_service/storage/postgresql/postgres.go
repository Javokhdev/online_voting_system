package postgresql

import (
	"database/sql"
	"fmt"
	"voting_service/config"
	"voting_service/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db          *sql.DB
	DbPub       *sql.DB
	ElectionS   storage.ElectionI
	CandidateS  storage.CandidateI
	PublicVoteS storage.PublicVoteI
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)

	connPub := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.PUB_DB_HOST, config.PUB_DB_USER, config.PUB_DB_NAME, config.PUB_DB_PASSWORD, config.PUB_DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	dbPub, err := sql.Open("postgres", connPub)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = dbPub.Ping()
	if err != nil {
		return nil, err
	}

	er := NewElectionStorage(db)
	cr := NewCandidateStorage(db, dbPub)
	pr := NewPublicVoteStorage(db, dbPub)

	return &Storage{
		Db:          db,
		DbPub:       dbPub,
		ElectionS:   er,
		CandidateS:  cr,
		PublicVoteS: pr,
	}, nil
}

func (s *Storage) Election() storage.ElectionI {
	if s.ElectionS == nil {
		s.ElectionS = NewElectionStorage(s.Db)
	}

	return s.ElectionS
}

func (s *Storage) Candidate() storage.CandidateI {
	if s.CandidateS == nil {
		s.CandidateS = NewCandidateStorage(s.Db, s.DbPub)
	}

	return s.CandidateS
}

func (s *Storage) PublicVote() storage.PublicVoteI {
	if s.PublicVoteS == nil {
		s.PublicVoteS = NewPublicVoteStorage(s.Db, s.DbPub)
	}

	return s.PublicVoteS
}
