package postgresql

import (
	"database/sql"
	v "voting_service/genproto/voting"

	"github.com/google/uuid"
)

type VoteStorage struct {
	db *sql.DB
}

func NewVoteStorage(db *sql.DB) *VoteStorage {
	return &VoteStorage{db: db}
}

func (s *VoteStorage) CreateVote(vote *v.CreateVoteReq) error {
	id := uuid.New().String()
	_, err := s.db.Exec("INSERT INTO votes (id, candidate_id) VALUES ($1, $2)",
		id, vote.GetCandidateId())
	return err
}

func (s *VoteStorage) GetVoteById(id *v.ById) (*v.GetVotesRes, error) {
	row := s.db.QueryRow("SELECT id, candidate_id FROM votes WHERE id = $1 and deleted_at = 0", id.GetId())
	votee := &v.GetVotesRes{}
	err := row.Scan(&votee.Id, &votee.CandidateId)
	return votee, err
}

func (s *VoteStorage) GetAllVotes(flt *v.Filter) (*v.GetAllVotesRes, error) {
	rows, err := s.db.Query("SELECT id, candidate_id FROM votes WHERE deleted_at = 0 LIMIT $1 OFFSET $2")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	votes := &v.GetAllVotesRes{}
	for rows.Next() {
		vote := &v.GetVotesRes{}
		err := rows.Scan(&vote.Id, &vote.CandidateId)
		if err != nil {
			return nil, err
		}
		votes.Votes = append(votes.Votes, vote)
	}
	return votes, nil
}

func (s *VoteStorage) UpdateVote(vote *v.GetVotesRes) (*v.Void, error) {
	_, err := s.db.Exec("UPDATE votes SET candidate_id = $1 WHERE id = $2",
		vote.GetCandidateId(), vote.GetId())
	return nil, err
}

func (s *VoteStorage) DeleteVote(id *v.ById) (*v.Void, error) {
	_, err := s.db.Exec("Update votes SET deleted = Extract(epoch from now()) WHERE id = $1", id.GetId())
	return nil, err
}