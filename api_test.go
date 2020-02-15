package IBank_core

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)


func TestLogin(t *testing.T) {
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
	_, err = db.Exec(usersDDL)
	if err != nil {
		t.Errorf("can't create table user err: %v", err)
	}

	err = AddUser(db, "don", "don", "don", "don", "don", false)
	if err != nil {
		t.Errorf("can't add user don: %v", err)
	}

	_, _, err = Login(db, "don", "xer")
	if err == nil {
		t.Errorf("Not Invalid Password Login() got err: %v", err)
	}
	_, _, err = Login(db, "don", "don")
	if err != nil {
		t.Errorf("Not correct pass Login() got err: %v", err)
	}
}

func TestAddUser(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("can't open db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf("can't close db: %v", err)
		}
	}()
	_, err = db.Exec(usersDDL)
	if err != nil {
		t.Errorf("can't create table user err: %v", err)
	}
	_, err = db.Exec(foreignKeysON)
	if err != nil {
		t.Errorf("can't on foreign keys: %v", err)
	}
	err = AddUser(db, "123", "123", "123", "123", "123", false)
	if err != nil {
		t.Errorf("can't insert user 123: %v", err)
	}
	err = AddUser(db, "123", "123", "123", "123", "123", false)
	if err == nil {
		t.Errorf("error user 123 already was but new user 123 was added: %v", err)
	}

}

func TestUsersList(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("can't open db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf("can't close db: %v", err)
		}
	}()
	_, err = db.Exec(usersDDL)
	if err != nil {
		t.Errorf("can't create table user err: %v", err)
	}
	_, err = db.Exec(foreignKeysON)
	if err != nil {
		t.Errorf("can't on foreign keys: %v", err)
	}
	userlist, err := UsersList(db)
	if got:=len(userlist); got != 0 {
		t.Errorf("no users but UserList() get got: %v, want: 0", got)
	}
	err = AddUser(db, "123", "123", "123", "123", "123", false)
	if err != nil {
		t.Errorf("can't insert user 123: %v", err)
	}
	err = AddUser(db, "1234", "1234", "1234", "1234", "1234", false)
	if err != nil {
		t.Errorf("can't insert user 124: %v", err)
	}
	list, err := UsersList(db)
	if err != nil {
		t.Errorf("can't get userlist err: %v", err)
	}
	want := []UserList{{Id: 1, Name: "123", Surname: "123", Phone: "123", Locked: false}, {Id: 2, Name: "1234", Surname: "1234", Phone: "1234", Locked: false}}
	if check(list[0], want[0]) && check(list[1], want[1]){

	}

}
func check(a,b UserList) bool {
	if a.Locked != b.Locked {
		return false
	}
	if a.Id != b.Id {
		return false
	}
	if a.Phone != b.Phone {
		return false
	}
	if a.Name != b.Name {
		return false
	}
	if a.Surname != b.Surname {
		return false
	}
	return true
}

func TestAddBillToUser(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("can't open db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf("can't close db: %v", err)
		}
	}()
	_, err = db.Exec(usersDDL)
	if err != nil {
		t.Errorf("can't create table users err: %v", err)
	}
	_, err = db.Exec(billsDDL)
	if err != nil {
		t.Errorf("can't create table bills err: %v", err)
	}
	_, err = db.Exec(foreignKeysON)
	if err != nil {
		t.Errorf("can't on foreign keys: %v", err)
	}
	err = AddUser(db, "123", "123", "123", "123", "123", false)
	if err != nil {
		t.Errorf("can't insert user 123: %v", err)
	}
	err = AddUser(db, "1234", "1234", "1234", "1234", "1234", true)
	if err != nil {
		t.Errorf("can't insert user 1234: %v", err)
	}
	err = AddBillToUser(db, 1, 100, false)
	if err != nil {
		t.Errorf("can't add bill to user 123: %v", err)
	}
	err = AddBillToUser(db, 2, 100, false)
	if err == nil {
		t.Errorf("user 1234 was blocked but AddBillToUser() got err: %v", err)
	}

}

func TestAddService(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("can't open db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf("can't close db: %v", err)
		}
	}()
	_, err = db.Exec(servicesDDL)
	if err != nil {
		t.Errorf("can't create table services err: %v", err)
	}
	_, err = db.Exec(foreignKeysON)
	if err != nil {
		t.Errorf("can't on foreign keys: %v", err)
	}
	err = AddService(db, "tjk", 200)
	if err != nil {
		t.Errorf("can't add service got err: %v", err)
	}
	err = AddService(db, "tjk", 200)
	if err == nil{
		t.Errorf("duplicate service added is no correct got err: %v", err)
	}

}

func TestAddATM(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("can't open db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf("can't close db: %v", err)
		}
	}()
	_, err = db.Exec(atmsDDL)
	if err != nil {
		t.Errorf("can't create table atms err: %v", err)
	}
	_, err = db.Exec(foreignKeysON)
	if err != nil {
		t.Errorf("can't on foreign keys: %v", err)
	}
	err = AddATM(db, "tjk", false)
	if err != nil {
		t.Errorf("can't add service got err: %v", err)
	}
	err = AddATM(db, "tjk", false)
	if err == nil{
		t.Errorf("duplicate service added is no correct got err: %v", err)
	}
}

func TestServicesList(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("can't open db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf("can't close db: %v", err)
		}
	}()
	_, err = db.Exec(servicesDDL)
	if err != nil {
		t.Errorf("can't create table services err: %v", err)
	}
	_, err = db.Exec(usersDDL)
	if err != nil {
		t.Errorf("can't create table users err: %v", err)
	}
	_, err = db.Exec(billsDDL)
	if err != nil {
		t.Errorf("can't create table bills err: %v", err)
	}
	_, err = db.Exec(foreignKeysON)
	if err != nil {
		t.Errorf("can't on foreign keys: %v", err)
	}

	err = AddUser(db,"123","123","123","123","123",false)
	if err != nil {
		t.Errorf("can't add user 123: %v", err)
	}
	err = AddUser(db,"1234","1234","1234","1234","1234",true)
	if err != nil {
		t.Errorf("can't add user 1234: %v", err)
	}
	err = AddBillToUser(db, 1, 100, false)
	if err != nil {
		t.Errorf("can't add bill to user 123: %v", err)
	}
	err = AddBillToUser(db, 2, 100, true)
	if err == nil {
		t.Errorf("can't add bill to user 123: %v", err)
	}
	_, err, _ = CheckBill(db, 1)
	if err != nil {
		t.Errorf("wrong err: %v", err)
	}
	_, err, _ = CheckBill(db, 2)
	if err == nil {
		t.Errorf("wrong err: %v", err)
	}
}

func TestCheckBill(t *testing.T) {

}

func TestAvailableBills(t *testing.T) {

}

func TestTransferBillToBill(t *testing.T) {

}

func Test_findUserByPhone(t *testing.T) {

}

func TestGetAnyBill(t *testing.T) {

}

func TestBillsWithUserList(t *testing.T) {

}