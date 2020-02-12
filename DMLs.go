package IBank_core

const BankDML = `insert into bank(name, balance) VALUES ("Alif", 0) on conflict do nothing;`