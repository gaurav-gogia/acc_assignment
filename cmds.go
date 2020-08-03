package main

import (
	"fmt"
)

func createCmd(nameflag, name, salflag, salary, ageflag, age string) error {
	if nameflag != NAMEFLAG || ageflag != AGEFLAG || salflag != SALARYFLAG {
		return ErrWrongArgument
	}
	return createEmp(name, salary, age)
}

func showCmd(empidflag, empid string) error {
	if empidflag != EMPIDFLAG {
		return ErrWrongArgument
	}
	return getOneEmp(empid)
}

func delCmd(empidflag, empid string) error {
	if empidflag != EMPIDFLAG {
		return ErrWrongArgument
	}
	return deleteEmp(empid)
}

func listCmd() error { return getAllEmp() }

func helpCmd() {
	fmt.Printf("Welcome to Accurics Tool \n\n")
	fmt.Println("USAGE: ")
	fmt.Printf("accurics <command> [options]\n\n")

	fmt.Println("COMMANDS & Options: ")
	fmt.Println("Please use one of the following commands")
	fmt.Println("\taccurics help -> To reach this menu")
	fmt.Println("\taccurics list -> To list all employees")
	fmt.Println("\taccurics create -name <name> -salary <salary> -age <age> -> To create a new employee")
	fmt.Println("\taccurics show -empid <empid> -> To get details of an employee")
	fmt.Println("\taccurics delete -empid <empid> -> To delete an employee")
}
