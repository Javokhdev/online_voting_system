package postgresql

import (
	"database/sql"
	v "voting_service/genproto/voting"

	"github.com/google/uuid"
)

type PublicVoteStorage struct {
	db *sql.DB
}

func NewPublicVoteStorage(db *sql.DB) *PublicVoteStorage {
	return &PublicVoteStorage{db: db}
}

func (s *PublicVoteStorage) CreatePublicVote(publicVote *v.CreatePublicVoteReq) (*v.Void, error) {
	id := uuid.New().String()
	_, err := s.db.Exec("INSERT INTO public_votes (id, election_id, public_id) VALUES ($1, $2, $3)",
	id, publicVote.GetElectionId(), publicVote.GetPublicId())
	return nil, err
}

func (s *PublicVoteStorage) GetPublicVoteById(id *v.ById) (*v.GetPublicVoteRes, error) {
	row := s.db.QueryRow("SELECT id, election_id, public_id FROM public_votes WHERE id = $1 and deleted_at = 0", id.GetId())
	publicVote := &v.GetPublicVoteRes{}
	err := row.Scan(&publicVote.Id, &publicVote.ElectionId, &publicVote.PublicId)
	return publicVote, err
}

func (s *PublicVoteStorage) GetAllPublicVotes(flt *v.Filter) (*v.GetAllPublicVoteRes, error) {
	rows, err := s.db.Query("SELECT id, election_id, public_id FROM public_votes WHERE deleted_at = 0 LIMIT $1 OFFSET $2", flt.GetLimit(), flt.GetOffset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	publicVotes := &v.GetAllPublicVoteRes{}

	for rows.Next() {
		publicVote := &v.GetPublicVoteRes{}
		err := rows.Scan(&publicVote.Id, &publicVote.ElectionId, &publicVote.PublicId)
		if err != nil {
			return nil, err
		}
		publicVotes.PublicVotes = append(publicVotes.PublicVotes, publicVote)
	}
	return publicVotes, nil
}

func (s *PublicVoteStorage) UpdatePublicVote(publicVote *v.GetPublicVoteRes) (*v.Void, error) {
	_, err := s.db.Exec("UPDATE public_votes SET election_id = $1, public_id = $2 WHERE id = $3 and deleted_at = 0",
		publicVote.GetElectionId(), publicVote.GetElectionId(), publicVote.GetId())
	return nil, err
}

func (s *PublicVoteStorage) DeletePublicVote(id *v.ById) (*v.Void, error) {
	_, err := s.db.Exec("Update public_votes SET deleted = Extract(epoch from now()) WHERE id = $1", id.GetId())
	return nil, err
}
