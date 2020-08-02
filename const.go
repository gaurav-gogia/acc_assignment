package main

import "fmt"

// Constant strings for all queries
const (
	INITDB = `
	create table if not exists employees (
		empid integer primary key autoincrement,
		name text not null,
		salary integer not null,
		age integer not null
	);
	`
	GETALL = "select * from employees;"
	GETONE = "select * from employees where empid = ?;"
	DELONE = "delete from employees where empid = ?;"
	SETONE = "insert into employees (name, salary, age) values(?, ?, ?);"
)

// Constant strings for all commands
const (
	HELP1 = "help"
	HELP2 = "h"

	LISTCMD   = "list"
	CREATECMD = "create"
	DELETECMD = "delete"
	SHOWCMD   = "show"
)

// Constant strings for all flags
const (
	HELP3 = "-help"
	HELP4 = "-h"

	NAMEFLAG   = "-name"
	SALARYFLAG = "-salary"
	AGEFLAG    = "-age"
	EMPIDFLAG  = "-empid"
)

// Constant strings for paths, driver, misc strings
const (
	DRIVER = "sqlite3"
	PATH   = "./dbdir/data.db"
)

// ErrNotEnoughArguments -> User defined error to explain it to the user tha they have not passed enough arguments // to the command line
var ErrNotEnoughArguments error = fmt.Errorf("Looks like some argument was missed. Please use command `accurics help` to get a better idea")

// ErrWrongArgument -> User defined error to explain it to the user tha they have not passed correct arguments
// to the command line
var ErrWrongArgument error = fmt.Errorf("You may have used wrong argument/flag. Please use command `accurics help` to get a better idea")
