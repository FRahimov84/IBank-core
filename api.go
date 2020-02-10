package IBank_core

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func Init(db *sql.DB) error {
	ddls := []string{usersDDL, billsDDL, atmsDDL, servicesDDL}
	for _, ddl := range ddls {
		_, err := db.Exec(ddl)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddUser(db *sql.DB, login, pass, name, surname, phone string, locked bool) error {
	_, err := db.Exec(`insert into users(login, pass, name, surname, phoneNumber, locked)
VALUES (?, ?, ?, ?, ?, ?);`, login, pass, name, surname, phone, locked)
	if err != nil {
		return fmt.Errorf("can't add a user in database %w", err)
	}
	return nil
}

func AddBillToUser(db *sql.DB, userId, balance int, locked bool) error {
	var userLock bool
	err := db.QueryRow(`select locked from users where id = ?;`, userId).Scan(&userLock)
	if err != nil {
		return fmt.Errorf("can't find a user with id: %v, %w", userId, err)
	}
	if userLock {
		return fmt.Errorf("this user blocked!!!")
	}
	_, err = db.Exec(`insert into bills(user_id, balance, locked)
VALUES (?, ?, ?);`, userId, balance, locked)
	if err != nil {
		return fmt.Errorf("can't add a bill to user_id: %v, %w", userId, err)
	}
	return nil
}

func AddService(db *sql.DB, service string, price int) error {
	_, err := db.Exec(`insert into services(name, price)
VALUES (?, ?);`, service, price)
	if err != nil {
		return fmt.Errorf("can't add a service: %v, %w", service, err)
	}
	return nil
}

func AddATM(db *sql.DB, address string, locked bool) error {
	_, err := db.Exec(`insert into ATMs(address, locked)
VALUES (?, ?);`, address, locked)
	if err != nil {
		return fmt.Errorf("can't add a ATM: %v, %w", address, err)
	}
	return nil
}

