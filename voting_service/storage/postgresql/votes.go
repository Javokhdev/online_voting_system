package postgresql

import (
	"database/sql"
	vote "voting_service/genproto/voting"

	"github.com/google/uuid"
)

type VoteStorage struct {
	db *sql.DB
}

func NewVoteStorage(db *sql.DB) *VoteStorage {
	return &VoteStorage{db: db}
}

func (s *VoteStorage) CreateVote(vote *vote.CreateVoteReq) error {
	id := uuid.New().String()
	_, err := s.db.Exec("INSERT INTO votes (id, candidate_id) VALUES ($1, $2)",
		id, vote.CandidateId)
	return err
}

func (s *VoteStorage) GetVoteById(id string) (*vote.GetVotesRes, error) {
	row := s.db.QueryRow("SELECT id, candidate_id FROM votes WHERE id = $1", id)
	votee := &vote.GetVotesRes{}
	err := row.Scan(&votee.Id, &votee.CandidateId)
	return votee, err
}

func (s *VoteStorage) GetAllVotes() ([]*vote.GetVotesRes, error) {
	rows, err := s.db.Query("SELECT id, candidate_id FROM votes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	votes := make([]*vote.GetVotesRes, 0)
	for rows.Next() {
		vote := &vote.GetVotesRes{}
		err := rows.Scan(&vote.Id, &vote.CandidateId)
		if err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}
	return votes, nil
}

func (s *VoteStorage) UpdateVote(vote *vote.GetVotesRes) error {
	_, err := s.db.Exec("UPDATE votes SET candidate_id = $2 WHERE id = $1",
		vote.Id, vote.CandidateId)
	return err
}

func (s *VoteStorage) DeleteVote(id string) error {
	_, err := s.db.Exec("Update votes SET deleted = Extract(epoch from now()) WHERE id = $1", id)
	return err
}

