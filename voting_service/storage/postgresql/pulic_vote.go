package postgresql

import (
	"database/sql"
	pvote "voting_service/genproto/voting"

	"github.com/google/uuid"
)

type PublicVoteStorage struct {
	db *sql.DB
}

func NewPublicVoteStorage(db *sql.DB) *PublicVoteStorage {
	return &PublicVoteStorage{db: db}
}

func (s *PublicVoteStorage) CreatePublicVote(publicVote *pvote.CreatePublicVoteReq) error {
	id := uuid.New().String()
	_, err := s.db.Exec("INSERT INTO public_votes (id, election_id, public_id) VALUES ($1, $2, $3)",
	id, publicVote.ElectionId, publicVote.PublicId)
	return err
}

func (s *PublicVoteStorage) GetPublicVoteById(id string) (*pvote.GetPublicVoteRes, error) {
	row := s.db.QueryRow("SELECT id, election_id, public_id FROM public_votes WHERE id = $1", id)
	publicVote := &pvote.GetPublicVoteRes{}
	err := row.Scan(&publicVote.Id, &publicVote.ElectionId, &publicVote.PublicId)
	return publicVote, err
}

func (s *PublicVoteStorage) GetAllPublicVotes() ([]*pvote.GetPublicVoteRes, error) {
	rows, err := s.db.Query("SELECT id, election_id, public_id FROM public_votes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	publicVotes := make([]*pvote.GetPublicVoteRes, 0)
	for rows.Next() {
		publicVote := &pvote.GetPublicVoteRes{}
		err := rows.Scan(&publicVote.Id, &publicVote.ElectionId, &publicVote.PublicId)
		if err != nil {
			return nil, err
		}
		publicVotes = append(publicVotes, publicVote)
	}
	return publicVotes, nil
}

func (s *PublicVoteStorage) UpdatePublicVote(publicVote *pvote.GetPublicVoteRes) error {
	_, err := s.db.Exec("UPDATE public_votes SET election_id = $2, public_id = $3 WHERE id = $1",
		publicVote.Id, publicVote.ElectionId, publicVote.PublicId)
	return err
}

func (s *PublicVoteStorage) DeletePublicVote(id string) error {
	_, err := s.db.Exec("Update public_votes SET deleted = Extract(epoch from now()) WHERE id = $1", id)
	return err
}
