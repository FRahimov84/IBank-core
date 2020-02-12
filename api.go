package IBank_core

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func Init(db *sql.DB) error {
	ddls := []string{foreignKeysON, usersDDL, billsDDL, atmsDDL, servicesDDL, BankDDL, BankDML}
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

func UsersList(db *sql.DB) ([]UserList, error) {
	rows, err := db.Query(`select id, name, surname, phoneNumber, locked from users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	usersList := make([]UserList, 0)
	for rows.Next() {
		userList := UserList{}
		err := rows.Scan(&userList.Id, &userList.Name, &userList.Surname, &userList.Phone, &userList.Locked)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, userList)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return usersList, nil
}

func ATMsList(db *sql.DB) ([]ATM, error) {
	rows, err := db.Query(`select id, address, locked from ATMs`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ATMsList := make([]ATM, 0)
	for rows.Next() {
		atm := ATM{}
		err := rows.Scan(&atm.Id, &atm.Address, &atm.Locked)
		if err != nil {
			return nil, err
		}
		ATMsList = append(ATMsList, atm)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ATMsList, nil
}

//  CLIENT

func Login(db *sql.DB, login, pass string) (int, string, error) {
	var user_id int
	var user_name string
	var user_pass string

	err := db.QueryRow(`select id, name, pass from users where login = ?`, login).Scan(&user_id, &user_name, &user_pass)
	if err != nil {
		return -1, "null", err
	}
	if user_pass != pass {
		return -1, "null", fmt.Errorf("invalid password")
	}
	return user_id, user_name, nil
}

func UserBills(db *sql.DB, user_id int) ([]BillList, error) {
	rows, err := db.Query(`select id, balance, locked from bills where user_id = ?`, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Bills := make([]BillList, 0)
	for rows.Next() {
		bill := BillList{}
		err := rows.Scan(&bill.Id, &bill.Balance, &bill.Locked)
		if err != nil {
			return nil, err
		}
		Bills = append(Bills, bill)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return Bills, nil
}

func PayService(db *sql.DB, service_id, user_id int) error {
	var dbServicePrice int
	err := db.QueryRow(`select price from services where id = ?`, service_id).Scan(&dbServicePrice)
	if err != nil {
		return fmt.Errorf("no such service id %w", err)
	}
	bills, err := AvailableBills(db, user_id, dbServicePrice)
	if err != nil {
		return err
	}
	for _, bill := range bills {
		fmt.Printf("%v\t%v\n", bill.Id, bill.Balance)
	}
	var chosed_id int
	fmt.Println("Введите id счета с которого оплатить:")
	_, err = fmt.Scan(&chosed_id)
	if err != nil {
		return fmt.Errorf("can't scan bill id %w", err)
	}
	for _, value := range bills {
		if value.Id == chosed_id {
			return payeer(db, value, dbServicePrice)
		}
	}
	return fmt.Errorf("No such bill id!")
}

func payeer(db *sql.DB, bill BillList, price int) error {
	var bankBalance int
	err := db.QueryRow(`select balance from Bank where id = 1`).Scan(&bankBalance)
	if err != nil {
		return fmt.Errorf("can't get Bank balance %w", err)
	}
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("can't open transaction %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	_, err = tx.Exec(`update bills set balance = ? where id = ?`, bill.Balance-price, bill.Id)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`update bank set balance = ? where id = 1`, bankBalance+price)
	if err != nil {
		return fmt.Errorf("can't change Bank balance %w", err)
	}
	return nil
}

func ServicesList(db *sql.DB) ([]Services, error) {
	rows, err := db.Query(`select id, name, price from services;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	servicesList := make([]Services, 0)
	for rows.Next() {
		service := Services{}
		err := rows.Scan(&service.Id, &service.Name, &service.Price)
		if err != nil {
			return nil, err
		}
		servicesList = append(servicesList, service)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return servicesList, nil
}

func CheckBill(db *sql.DB, bill_id int) (bool, error, int) {
	var lock bool
	var balance int
	err := db.QueryRow(`select locked, balance from bills where id = ?`, bill_id).Scan(&lock, &balance)
	if err != nil {
		return false, err, -1
	}
	if lock {
		return false, fmt.Errorf("This bill is blocked!"), -1
	}
	return true, nil, balance
}

func AvailableBills(db *sql.DB, user_id int, amount int) ([]BillList, error) {
	bills, err := UserBills(db, user_id)
	if err != nil {
		return nil, err
	}
	solvensy := false
	availableBills := make([]BillList, 0)
	for _, value := range bills {
		if value.Balance >= amount && !value.Locked {
			availableBills = append(availableBills, value)
			solvensy = true
		}
	}
	if !solvensy {
		return nil, fmt.Errorf("you can't pay this amount")
	}
	return availableBills, nil
}

func TransferBillToBill(db *sql.DB, sender_id, sender_balance, addressee_id, addressee_balance, amount int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("can't open transaction %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	_, err = tx.Exec(`update bills set balance = ? where id = ?`, sender_balance-amount, sender_id)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`update bills set balance = ? where id = ?`, addressee_balance+amount, addressee_id)
	if err != nil {
		return err
	}
	return nil
}

func findUserByPhone(db *sql.DB, phone string) (int, error){
	var addressee_id int
	err := db.QueryRow(`select id from users where phoneNumber = ?`, phone).Scan(&addressee_id)
	if err != nil {
		return -1, err
	}
	return addressee_id, nil
}

func GetAnyBill(db *sql.DB, phone string, amount int) (int, int, error) {
	addressee_id, err := findUserByPhone(db, phone)
	if err != nil {
		return -1,-1,err
	}
	bills, err := AvailableBills(db, addressee_id, amount)
	if err != nil {
		return -1, -1, err
	}
	return bills[0].Id, bills[0].Balance, nil
}