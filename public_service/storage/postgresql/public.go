package postgresql	

import (
	"database/sql"
	"github.com/google/uuid"
	p "public_service/genproto/public"
)

type PublicSotage struct {
	Db *sql.DB
}

func NewPublicStorage(db *sql.DB) *PublicSotage {
	return &PublicSotage{Db: db}
}

func (ps *PublicSotage) Create(public *p.CreatePublicReq) (*p.Void, error) {
	id := uuid.New().String()
	_, err := ps.Db.Exec("INSERT INTO publics(id, first_name, last_name, birthday, gender, nation, party_id) VALUES($1, $2, $3, $4, $5, $6, $7)",
	id, public.GetFirstName(), public.GetLastName(), public.GetBirthday(), public.GetGender(), public.GetNation(), public.GetPartyId())
	return nil, err	
}

func (ps *PublicSotage) GetById(id *p.ById) (*p.GetPublicRes, error) {
	public := &p.GetPublicRes{}
	row := ps.Db.QueryRow("SELECT id, first_name, last_name, birthday, gender, nation, party_id FROM publics WHERE id = $1", id.GetId())
	err := row.Scan(public.id, &public.FirstName, &public.LastName, &public.Birthday, &public.Gender, &public.Nation, &public.PartyId)
	if err != nil {
		return nil, err
	}
	return public, nil
}

func (ps *PublicSotage) GetAll(flt *p.Filter) (*p.GetAllPublicRes, error) {
	rows, err := ps.Db.Query("SELECT id, first_name, last_name, birthday, gender, nation, party_id FROM publics where deleted_at = 0 LIMIT $1 OFFSET $2", flt.GetLimit(), flt.GetOffset())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	publics := &p.GetAllPublicRes{}
	for rows.Next() {
		public := &p.GetPublicRes{}
		if err := rows.Scan(&public.Id, &public.FirstName, &public.LastName, &public.Birthday, &public.Gender, &public.Nation, &public.PartyId); err != nil {
			return nil, err
		}
		publics.Parties = append(publics.Parties, public)
	}
	return publics, nil
}

func (ps *PublicSotage) Update(public *p.GetPublicRes) (*p.Void, error) {
	_, err := ps.Db.Exec("UPDATE publics SET first_name = $1, last_name = $2, birthday = $3, gender = $4, nation = $5, party_id = $6 WHERE id = $7",
	public.GetFirstName(), public.GetLastName(), public.GetBirthday(), public.GetGender(), public.GetNation(), public.GetPartyId(), public.GetId())
	return nil, err
}

func (ps *PublicSotage) Delete(id *p.ById) (*p.Void, error) {
	_, err := ps.Db.Exec("Update publics SET deleted_at = Extract(epoch from NOW()) WHERE id = $1", id.GetId())
	if err != nil {
		return nil, err
	}

	return nil, err
}