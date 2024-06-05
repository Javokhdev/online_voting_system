package postgresql

import (
	"database/sql"
)

type CandidateStorage struct {
	db *sql.DB
}

func NewCandidateStorage(db *sql.DB) *CandidateStorage {
	return &CandidateStorage{db: db}
}
