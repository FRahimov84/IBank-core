package IBank_core

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