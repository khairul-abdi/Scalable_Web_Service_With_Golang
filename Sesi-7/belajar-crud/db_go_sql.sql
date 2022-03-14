create database db_go_sql;

CREATE TABLE public.employees (
	id SERIAL primary key,
	full_name varchar(50) NOT NULL,
	email text unique not null,
	age int not null,
	division varchar(20) not null
);