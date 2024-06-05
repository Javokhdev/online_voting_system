package postgresql

import (
	"database/sql"
	"github.com/google/uuid"
	elec "voting_service/genproto/election"
)

type ElectionStorage struct {
	db *sql.DB
}

func NewElectionStorage(db *sql.DB) *ElectionStorage {
	return &ElectionStorage{db: db}
}

func (s *ElectionStorage) CreateElection(election *elec.CreateElectionReq) error {
	id := uuid.New().String()
	_, err := s.db.Exec("INSERT INTO elections (id, name, date) VALUES ($1, $2, $3)", 
	id, election.Name, election.Date)

	return err
}

func (s *ElectionStorage) GetElectionById(id string) (*elec.GetElectionRes, error) {
	row := s.db.QueryRow("SELECT id, name, date FROM elections WHERE id = $1", id)
	election := &elec.GetElectionRes{}
	err := row.Scan(&election.Id, &election.Name, &election.Date)
	return election, err
}

func (s *ElectionStorage) GetAllElections() ([]*elec.GetElectionRes, error) {
	rows, err := s.db.Query("SELECT id, name, date FROM elections")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	elections := make([]*elec.GetElectionRes, 0)
	for rows.Next() {
		election := &elec.GetElectionRes{}
		err := rows.Scan(&election.Id, &election.Name, &election.Date)
		if err != nil {
			return nil, err
		}
		elections = append(elections, election)
	}
	return elections, nil
}

func (s *ElectionStorage) UpdateElection(election *elec.GetElectionRes) error {
	_, err := s.db.Exec("UPDATE elections SET name = $2, date = $3 WHERE id = $1",
		election.Id, election.Name, election.Date)
	return err
}

func (s *ElectionStorage) DeleteElection(id string) error {
	_, err := s.db.Exec("Update elections SET deleted = Extract(epoch from now()) WHERE id = $1", id)
	return err
}