package main

// AllEmployees structure will be used to deserealize json string response from /employee API into golang structure
type AllEmployees struct {
	Status string       `json:"status"`
	Data   []allEmpData `json:"data"`
}

// OneEmployee structure will be used to deserealize json string response from /employee/{} & /create API into golang structure
type OneEmployee struct {
	Status string     `json:"status"`
	Data   oneEmpData `json:"data"`
}

// ModifyEmployee structure will be used to deserealize json string response from /delete/{} API into golang structure
type ModifyEmployee struct {
	Status  string `json:"status"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

// oneEmpData is the supporting data structure for OneEmployee structure
type oneEmpData struct {
	ID             int    `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
	EmployeeAge    int    `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

// allEmpData is the supporting data structure for AllEmployees structure
type allEmpData struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}
