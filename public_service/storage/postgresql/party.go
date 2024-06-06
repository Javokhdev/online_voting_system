package postgresql

import (
	"database/sql"
	"github.com/google/uuid"
	p "public_service/genproto/public"
)

type PartyStorage struct {
	Db *sql.DB
}

func NewPartyStorage(db *sql.DB) *PartyStorage {
	return &PartyStorage{Db: db}
}

func (ps *PartyStorage) Create(party *p.CreatePartyReq) (*p.Void, error) {
	id := uuid.New().String()
	_, err := ps.Db.Exec("INSERT INTO parties(id, name, slogan, opened_date, description) VALUES($1, $2, $3, $4, $5)",
	id, party.GetName(), party.GetSlogan(), party.GetOpenedDate(), party.GetDescription())

	return nil, err	
}

func (ps *PartyStorage) GetById(id *p.ById) (*p.GetPartyRes, error) {
	party := &p.GetPartyRes{}
	err := ps.Db.QueryRow("SELECT id, name, slogan, opened_date, description FROM parties WHERE id = $1 and deleted_at = 0", id.GetId()).Scan(&party.Id, &party.Name, &party.Slogan, &party.OpenedDate, &party.Description)
	if err != nil {
		return nil, err
	}

	return party, err
}
 
func (ps *PartyStorage) GetAll(flt *p.Filter) (*p.GetAllPartyRes, error) {
	rows, err := ps.Db.Query("SELECT id, name, slogan, opened_date, description FROM parties where deleted_at = 0 LIMIT $1 OFFSET $2", flt.GetLimit(), flt.GetOffset())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	parties := &p.GetAllPartyRes{}

	for rows.Next() {
		party := &p.GetPartyRes{}

		if err := rows.Scan(&party.Id, &party.Name, &party.Slogan, &party.OpenedDate, &party.Description); err != nil {
			return nil, err
		}

		parties.Parties = append(parties.Parties, party)
	}

	return parties, nil
}

func (ps *PartyStorage) Update(party *p.GetPartyRes) (*p.Void, error) {
	_, err := ps.Db.Exec("UPDATE parties SET name = $1, slogan = $2, opened_date = $3, description = $4, updated_at = now() WHERE id = $5 and deleted_at = 0",
		party.GetName(), party.GetSlogan(), party.GetOpenedDate(), party.GetDescription(), party.GetId())
	return nil, err
}

func (ps *PartyStorage) Delete(id *p.ById) (*p.Void, error) {
	_, err := ps.Db.Exec("Update parties SET deleted_at = Extarct(EPOCH FROM NOW()) WHERE id = $1 and deleted_at = 0", id.GetId())
	
	return nil, err
}

