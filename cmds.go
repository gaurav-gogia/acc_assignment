package main

import (
	"database/sql"
	"fmt"
)

func createCmd(db *sql.DB, nameflag, name, ageflag, age, salflag, salary string) error {
	if nameflag != NAMEFLAG || ageflag != AGEFLAG || salflag != SALARYFLAG {
		return ErrWrongArgument
	}
	return createEmp(db, name, age, salary)
}

func showCmd(db *sql.DB, empidflag, empid string) error {
	if empidflag != EMPIDFLAG {
		return ErrWrongArgument
	}
	return showEmp(db, empid)
}

func delCmd(db *sql.DB, empidflag, empid string) error {
	if empidflag != EMPIDFLAG {
		return ErrWrongArgument
	}
	return delEmp(db, empid)
}

func listCmd(db *sql.DB) error { return showEmp(db, "") }

func helpCmd() {
	fmt.Printf("Welcome to Accurics Tool \n\n")
	fmt.Println("USAGE: ")
	fmt.Printf("accurics <command> [options]\n\n")

	fmt.Println("COMMANDS & Options: ")
	fmt.Println("Please use one of the following commands")
	fmt.Println("\taccurics help -> To reach this menu")
	fmt.Println("\taccurics list -> To list all employees")
	fmt.Println("\taccurics create -name <name> -age <age> -salary <salary> -> To create a new employee")
	fmt.Println("\taccurics show -empid <empid> -> To get details of an employee")
	fmt.Println("\taccurics delete -empid <empid> -> To delete an employee")
}
