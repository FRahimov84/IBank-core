package IBank_core

type List struct {
	UsersList []UserList
	ATMsList []ATM
	BillUserList []BillUser
}

type UserList struct {
	Id int
	Name string
	Surname string
	Phone string
	Locked bool
}

type ATM struct {
	Id int
	Address string
	Locked bool
}

type Services struct {
	Id int
	Name string
	Price int
}

type BillList struct{
	Id int
	Balance int
	Locked bool
}
type BillUser struct {
	Id int
	Balance int
	LockedBill bool
	UserName string
	UserSurname string
	UserPhone string
	LockedUser bool
}