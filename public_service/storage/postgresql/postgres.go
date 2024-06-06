package postgresql

import (
	"database/sql"
	"fmt"
	"public_service/config"
	"public_service/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db          *sql.DB
	PartyS   storage.PartyI
	PublicS  storage.PublicI
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

	pr := NewPartyStorage(db)
	pb := NewPublicStorage(db)

	return &Storage{
		Db:          db,
		PartyS:   pr,
		PublicS:  pb,
	}, nil
}

func (s *Storage) Party()storage.PartyI{
	if s.PartyS == nil {
		s.PartyS = NewPartyStorage(s.Db)
	}

	return s.PartyS
}

func (s *Storage) Public()storage.PublicI{
	if s.PublicS == nil {
		s.PublicS = NewPublicStorage(s.Db)
	}

	return s.PublicS
}