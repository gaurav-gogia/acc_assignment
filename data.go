package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func initdb() (*sql.DB, error) {
	db, err := sql.Open(DRIVER, PATH)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(INITDB)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec()

	return db, err
}

func createEmp(db *sql.DB, name, salary, age string) error {
	stmt, err := db.Prepare(SETONE)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, salary, age)
	if err != nil {
		return err
	}

	fmt.Println("Employee Created")
	return nil
}

func delEmp(db *sql.DB, empid string) error {
	stmt, err := db.Prepare(DELONE)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(empid)
	if err != nil {
		return err
	}

	fmt.Println("Employee Deleted")
	return nil
}

func showEmp(db *sql.DB, empid string) error {
	var queryStr string
	if empid == "" {
		queryStr = GETALL
	} else {
		queryStr = GETONE
	}

	var id, name, salary, age string
	rows, err := db.Query(queryStr, empid)
	if err != nil {
		return err
	}

	for rows.Next() {
		rows.Scan(&id, &name, &salary, &age)
		fmt.Printf("Employee ID: %s\n", id)
		fmt.Println("Employee Name: ", name)
		fmt.Println("Employee Salary: ", salary)
		fmt.Println("Employee Age: ", age)
		fmt.Printf("###############################################################\n")
	}

	return nil
}
