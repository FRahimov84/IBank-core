package IBank_core

const usersDDL = `create table if not exists users
(
    id          integer primary key autoincrement,
    login       text unique not null,
    pass        text        not null,
    name        text        not null,
    surname     text        not null,
    phoneNumber text unique not null,
    locked      boolean     not null
);`
const billsDDL = `create table if not exists bills
(
    id      integer primary key autoincrement,
    user_id integer references users,
    balance integer not null check ( balance > 0 ),
    locked  boolean not null
);`
const ATMsDDL = `create table if not exists ATMs
(
    id      integer primary key autoincrement,
    address text    not null,
    locked  boolean not null
);`
const servicesDDL = `create table if not exists services
(
    id     integer primary key autoincrement,
    name   text    not null unique,
    price integer not null check ( price > 0 )
);`