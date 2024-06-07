package postgresql

import (
	"database/sql"
	"errors"
	v "voting_service/genproto/voting"

	"github.com/google/uuid"
)

type PublicVoteStorage struct {
	db    *sql.DB
	dbPub *sql.DB
}

func NewPublicVoteStorage(db *sql.DB, dbPub *sql.DB) *PublicVoteStorage {
	return &PublicVoteStorage{db: db, dbPub: dbPub}
}

func (s *PublicVoteStorage) Create(publicVote *v.CreatePublicVoteReq) (*v.Void, error) {
	tz, err := s.db.Begin()

	defer func() {
		tz.Commit()
	}()
	defer func() {
		if err != nil {
			err = tz.Rollback()
			if err != nil {
				panic(err)
			}
		}
	}()

	id := uuid.New().String()

	if !(s.CheckPublic(publicVote.PublicId)) {
		err := errors.New("error while creating")
		return nil, err
	}

	_, err = tz.Exec("INSERT INTO public_votes (id, election_id, public_id) VALUES ($1, $2, $3)",
		id, publicVote.GetElectionId(), publicVote.GetPublicId())

	id = uuid.New().String()
	_, err = tz.Exec("INSERT INTO votes (id, candidate_id) VALUES ($1, $2)",
		id, publicVote.GetCandidateId())

	return nil, err
}

func (s *PublicVoteStorage) GetById(id *v.ById) (*v.GetPublicVoteRes, error) {
	row := s.db.QueryRow("SELECT id, election_id, public_id FROM public_votes WHERE id = $1 and deleted_at = 0", id.GetId())
	publicVote := &v.GetPublicVoteRes{}
	err := row.Scan(&publicVote.Id, &publicVote.ElectionId, &publicVote.PublicId)
	return publicVote, err
}

func (s *PublicVoteStorage) GetAll(flt *v.Filter) (*v.GetAllPublicVoteRes, error) {
	rows, err := s.db.Query("SELECT id, election_id, public_id FROM public_votes WHERE deleted_at = 0 LIMIT $1 OFFSET $2", flt.GetLimit(), flt.GetOffset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	publicVotes := &v.GetAllPublicVoteRes{}

	for rows.Next() {
		publicVote := &v.GetPublicVoteRes{}
		err := rows.Scan(&publicVote.Id, &publicVote.ElectionId, &publicVote.PublicId)
		if err != nil {
			return nil, err
		}
		publicVotes.PublicVotes = append(publicVotes.PublicVotes, publicVote)
	}
	return publicVotes, nil
}

func (s *PublicVoteStorage) Update(publicVote *v.GetPublicVoteRes) (*v.Void, error) {
	_, err := s.db.Exec("UPDATE public_votes SET election_id = $1, public_id = $2 WHERE id = $3 and deleted_at = 0",
		publicVote.GetElectionId(), publicVote.GetElectionId(), publicVote.GetId())
	return nil, err
}

func (s *PublicVoteStorage) Delete(id *v.ById) (*v.Void, error) {
	_, err := s.db.Exec("Update public_votes SET deleted = Extract(epoch from now()) WHERE id = $1", id.GetId())
	return nil, err
}

func (s *PublicVoteStorage) GetVote(vote *v.ById) (*v.GetVoteById, error) {

	row, err := s.db.Query("select candidate_id, count(id) from votes where candidate_id = $1 group by candidate_id", vote.GetId())
	if err != nil {
		return nil, err
	}

	defer row.Close()

	res := &v.GetVoteById{}

	err = row.Scan(&res.CandidateId, &res.Count)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PublicVoteStorage) GetVotes(flt *v.Filter) (*v.GetAllVotes, error) {

	rows, err := s.db.Query("select candidate_id, count(id) from votes group by candidate_id LIMIT $1 OFFSET $2", flt.GetLimit(), flt.GetOffset())

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	votes := &v.GetAllVotes{}

	for rows.Next() {
		vote := &v.GetVoteById{}
		err := rows.Scan(&vote.CandidateId, &vote.Count)
		if err != nil {
			return nil, err
		}
		votes.Votes = append(votes.Votes, vote)
	}
	return votes, nil
}

func (s *PublicVoteStorage) CheckPublic(id string) bool {

	var findId string

	row := s.dbPub.QueryRow("SELECT id FROM publics WHERE id = $1 and deleted_at = 0", id)

	err := row.Scan(&findId)

	if err != nil {
		return false
	}

	return findId == id
}
