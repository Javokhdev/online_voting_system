package postgresql

import (
	"database/sql"
	"fmt"
	"voting_service/config"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db          *sql.DB
	ElectionS   *ElectionStorage
	CandidateS  *CandidateStorage
	PublicVoteS *PublicVoteStorage
	VotesS      *VoteStorage
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	er := NewElectionStorage(db)
	cr := NewCandidateStorage(db)
	pr := NewPublicVoteStorage(db)
	vr := NewVoteStorage(db)

	return &Storage{
		Db:          db,
		ElectionS:   er,
		CandidateS:  cr,
		PublicVoteS: pr,
		VotesS:      vr,
	}, nil
}
