package postgresql

import (
	"database/sql"
	v "voting_service/genproto/voting"

	"github.com/google/uuid"
)

type CandidateStorage struct {
	db    *sql.DB
}

func NewCandidateStorage(db *sql.DB) *CandidateStorage {
	return &CandidateStorage{db: db}
}

func (s *CandidateStorage) Create(candidate *v.CreateCandidateReq) (*v.Void, error) {
	id := uuid.New().String()

	_, err := s.db.Exec("INSERT INTO candidates (id, election_id, public_id, party_id) VALUES ($1, $2, $3, $4)",
		id, candidate.GetElectionId(), candidate.GetPublicId(), candidate.GetPartyId())

	return nil, err
}

func (s *CandidateStorage) GetById(id *v.ById) (*v.GetCandidateRes, error) {
	row := s.db.QueryRow("SELECT id, election_id, public_id, party_id FROM candidates WHERE id = $1", id.GetId())
	candidate := &v.GetCandidateRes{}
	err := row.Scan(&candidate.Id, &candidate.ElectionId, &candidate.PublicId, &candidate.PartyId)
	if err != nil {
		return nil, err
	}
	return candidate, err
}

func (s *CandidateStorage) GetAll(flt *v.Filter) (*v.GetAllCandidateRes, error) {
	rows, err := s.db.Query("SELECT id, election_id, public_id, party_id FROM candidates where deleted_at = 0 LIMIT $1 OFFSET $2", flt.GetLimit(), flt.GetOffset())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	candidates := &v.GetAllCandidateRes{}
	for rows.Next() {
		candidate := &v.GetCandidateRes{}
		err := rows.Scan(&candidate.Id, &candidate.ElectionId, &candidate.PublicId, &candidate.PartyId)
		if err != nil {
			return nil, err
		}
		candidates.Candidates = append(candidates.Candidates, candidate)
	}
	return candidates, nil
}

func (s *CandidateStorage) Update(candidate *v.GetCandidateRes) (*v.Void, error) {
	_, err := s.db.Exec("UPDATE candidates SET election_id = $1, public_id = $2, party_id = $3 WHERE id = $4 and deleted_at = 0",
		candidate.GetElectionId(), candidate.GetPublicId(), candidate.GetPartyId(), candidate.GetId())

	return nil, err
}

func (s *CandidateStorage) Delete(id *v.ById) (*v.Void, error) {
	_, err := s.db.Exec("Update candidates SET deleted = Extract(epoch from now()) WHERE id = $1 and deleted_at = 0", id.GetId())
	return nil, err
}