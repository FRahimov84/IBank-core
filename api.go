package IBank_core

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Init(db *sql.DB) error{
	ddls := []string{usersDDL, billsDDL, ATMsDDL, servicesDDL}
	for _,ddl := range ddls{
		_, err := db.Exec(ddl)
		if err != nil {
			return err
		}
	}
	return nil
}

