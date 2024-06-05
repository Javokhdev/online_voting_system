package postgresql

import (
	"database/sql"
	"github.com/google/uuid"
	cand "voting_service/genproto/candidate"
)

type CandidateStorage struct {
	db *sql.DB
}

func NewCandidateStorage(db *sql.DB) *CandidateStorage {
	return &CandidateStorage{db: db}
}

func (s *CandidateStorage) CreateCandidate(candidate *cand.CreateCandidateReq) error {
	id := uuid.New().String()
	_, err := s.db.Exec("INSERT INTO candidates (id, election_id, public_id, party_id) VALUES ($1, $2, $3, $4)", 
	id, candidate.ElectionId, candidate.PublicId, candidate.PartyId)
	if err != nil {
		return err
	}


	return err
}

func (s *CandidateStorage) GetCandidateById(id string) (*cand.GetCandidateRes, error) {
	row := s.db.QueryRow("SELECT id, election_id, public_id, party_id FROM candidates WHERE id = $1", id)
	candidate := &cand.GetCandidateRes{}
	err := row.Scan(&candidate.Id, &candidate.ElectionId, &candidate.PublicId, &candidate.PartyId)
	if err != nil {
		return nil, err
	}
	return candidate, err
}

func (s *CandidateStorage) GetAllCandidates() ([]*cand.GetCandidateRes, error) {
	rows, err := s.db.Query("SELECT id, election_id, public_id, party_id FROM candidates")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	candidates := make([]*cand.GetCandidateRes, 0)
	for rows.Next() {
		candidate := &cand.GetCandidateRes{}
		err := rows.Scan(&candidate.Id, &candidate.ElectionId, &candidate.PublicId, &candidate.PartyId)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, candidate)
	}
	return candidates, nil
}

func (s *CandidateStorage) UpdateCandidate(candidate *cand.GetCandidateRes) error {
	_, err := s.db.Exec("UPDATE candidates SET election_id = $2, public_id = $3, party_id = $4 WHERE id = $1",
		candidate.Id, candidate.ElectionId, candidate.PublicId, candidate.PartyId)
	return err
}

func (s *CandidateStorage) DeleteCandidate(id string) error {
	_, err := s.db.Exec("Update candidates SET deleted = Extract(epoch from now()) WHERE id = $1", id)
	return err
}