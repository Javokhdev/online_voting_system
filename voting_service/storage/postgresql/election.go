package postgresql

import (
	"database/sql"
	"github.com/google/uuid"
	v "voting_service/genproto/voting"
)

type ElectionStorage struct {
	db *sql.DB
}

func NewElectionStorage(db *sql.DB) *ElectionStorage {
	return &ElectionStorage{db: db}
}

func (s *ElectionStorage) CreateElection(election *v.CreateElectionReq) (*v.Void, error) {
	id := uuid.New().String()
	_, err := s.db.Exec("INSERT INTO elections (id, name, date) VALUES ($1, $2, $3)", 
	id, election.Name, election.Date)

	return nil, err
}

func (s *ElectionStorage) GetElectionById(id *v.ById) (*v.GetElectionRes, error) {
	row := s.db.QueryRow("SELECT id, name, date FROM elections WHERE id = $1 and deleted_at = 0", id.GetId())
	election := &v.GetElectionRes{}
	err := row.Scan(&election.Id, &election.Name, &election.Date)
	return election, err
}

func (s *ElectionStorage) GetAllElections(flt *v.Filter) (*v.GetAllElectionRes, error) {
	rows, err := s.db.Query("SELECT id, name, date FROM elections WHERE deleted_at = 0 LIMIT $1 OFFSET $2", flt.GetLimit(), flt.GetOffset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	elections := &v.GetAllElectionRes{}
	for rows.Next() {
		election := &v.GetElectionRes{}
		err := rows.Scan(&election.Id, &election.Name, &election.Date)
		if err != nil {
			return nil, err
		}
		elections.Elections = append(elections.Elections, election)
	}
	return elections, nil
}

func (s *ElectionStorage) UpdateElection(election *v.GetElectionRes) (*v.Void, error) {
	_, err := s.db.Exec("UPDATE elections SET name = $1, date = $2 WHERE id = $3 and deleted_at = 0",
	election.GetName(), election.GetDate(), election.GetId())

	return nil, err
}

func (s *ElectionStorage) DeleteElection(id *v.ById) (*v.Void, error) {
	_, err := s.db.Exec("Update elections SET deleted_at = Extract(epoch from now()) WHERE id = $1", id.GetId())

	return nil, err
}