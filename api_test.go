package IBank_core

import (
	"database/sql"
	"testing"
)


func TestLoginClient_LoginNotOkForInvalidPassword(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("can't open db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf("can't close db: %v", err)
		}
	}()

	// shift 2 раза -> sql dialect
	_, err = db.Exec(`
  CREATE TABLE clients (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
  login TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL)`)
	if err != nil {
		t.Errorf("can't execute Login: %v", err)
	}

	_, err = db.Exec(`INSERT INTO clients(id, login, password) VALUES (1, 'don', 'don');`)
	if err != nil {
		t.Errorf("can't execute Login: %v", err)
	}

	_, _, err = Login(db, "don", "xer")

	if err == nil {
		t.Errorf("Not ErrInvalidPass error for invalid pass: %v", err)
	}
}
