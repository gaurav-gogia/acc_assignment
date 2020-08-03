package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func createEmp(name, salary, age string) error {
	var empData ModifyEmployee
	postData := fmt.Sprintf(`
		{
			"name": "%s",
			"salary": "%s",
			"age": "%s"
		}
	`, name, salary, age)

	res, err := http.Post(CREATEURL, "application/json", strings.NewReader(postData))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ErrAPICall
	}

	err = json.NewDecoder(res.Body).Decode(&empData)
	if err != nil {
		return err
	}

	if empData.Status != SUCCESS {
		return ErrAPICall
	}

	fmt.Println(empData.Message)
	return nil
}

func getAllEmp() error {
	var allEmpData AllEmployees

	res, err := http.Get(GETALLURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ErrAPICall
	}

	err = json.NewDecoder(res.Body).Decode(&allEmpData)
	if err != nil {
		return err
	}

	if allEmpData.Status != SUCCESS {
		return ErrAPICall
	}

	for _, emp := range allEmpData.Data {
		fmt.Println("Employee ID: ", emp.ID)
		fmt.Println("Employee Name: ", emp.EmployeeName)
		fmt.Println("Employee Salary: ", emp.EmployeeSalary)
		fmt.Println("Employee Age: ", emp.EmployeeAge)
		fmt.Println("################################################################")
	}

	return nil
}

func getOneEmp(empid string) error {
	var empData OneEmployee

	res, err := http.Get(GETONEURL + empid)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ErrAPICall
	}

	err = json.NewDecoder(res.Body).Decode(&empData)
	if err != nil {
		return err
	}

	if empData.Status != SUCCESS {
		return ErrAPICall
	}

	fmt.Println("Employee ID: ", empData.Data.ID)
	fmt.Println("Employee Name: ", empData.Data.EmployeeName)
	fmt.Println("Employee Salary: ", empData.Data.EmployeeSalary)
	fmt.Println("Employee Age: ", empData.Data.EmployeeAge)

	return nil
}

func deleteEmp(empid string) error {
	var empData ModifyEmployee

	var client http.Client
	req, err := http.NewRequest(http.MethodDelete, DELETEURL+empid, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ErrAPICall
	}

	err = json.NewDecoder(res.Body).Decode(&empData)
	if err != nil {
		return err
	}

	if empData.Status != SUCCESS {
		return ErrAPICall
	}

	fmt.Println(empData.Message)
	return nil
}
