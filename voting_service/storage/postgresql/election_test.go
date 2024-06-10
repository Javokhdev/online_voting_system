package postgresql

import (
	"database/sql"
	"testing"
	"time"
	v "voting_service/genproto/voting"
)

func TestElectionStorage(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=Qodir1111 dbname=voting sslmode=disable")
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	storage := NewElectionStorage(db)

	// Test Create method
	createReq := &v.CreateElectionReq{Name: "2023 City Council", Date: time.Now().Format("2024-01-01")}
	_, err = storage.Create(createReq)
	if err != nil {
		t.Errorf("failed to create election: %v", err)
	}

	// // Test GetById method
	getByID := &v.ById{Id: "55555555-5555-5555-5555-555555555555"}
	election, err := storage.GetById(getByID)
	if err != nil {
		if err == sql.ErrNoRows {
			t.Errorf("no rows found")
		}
		t.Errorf("failed to get election by id: %v", err)
	}
	if election.GetName() != createReq.GetName() {
		t.Errorf("expected election name %s, got %s", createReq.GetName(), election.GetName())
	}

	// Test GetAll method
	getAllReq := &v.Filter{Limit: 10, Offset: 0}
	elections, err := storage.GetAll(getAllReq)
	if err != nil {
		t.Errorf("failed to get all elections: %v", err)
	}
	if len(elections.GetElections()) != 10 {
		t.Errorf("expected 1 election, got %d", len(elections.GetElections()))
	}

	// Test Update method
	updateReq := &v.GetElectionRes{Id: "09293af0-a020-4974-a13c-6e6f7f6dce8a", Name: "Updated Election", Date: time.Now().Format("2006-01-02")}
	_, err = storage.Update(updateReq)
	if err != nil {
		t.Errorf("failed to update election: %v", err)
	}

	// Test Delete method
	deleteReq := &v.ById{Id: "09293af0-a020-4974-a13c-6e6f7f6dce8a"}
	_, err = storage.Delete(deleteReq)
	if err != nil {
		t.Errorf("failed to delete election: %v", err)
	}
}
