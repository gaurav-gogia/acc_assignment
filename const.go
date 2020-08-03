package main

import "fmt"

// SUCCESS string constant for response
const SUCCESS = "success"

// Constant strings for APIs
const (
	BASEURL   = "http://dummy.restapiexample.com/api/v1/"
	GETALLURL = BASEURL + "employees"
	GETONEURL = BASEURL + "employee/"
	CREATEURL = BASEURL + "create"
	DELETEURL = BASEURL + "delete/"
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

// ErrNotEnoughArguments -> User defined error to explain it to the user tha they have not passed enough arguments // to the command line
var ErrNotEnoughArguments error = fmt.Errorf("Looks like some argument was missed. Please use command `accurics help` to get a better idea")

// ErrWrongArgument -> User defined error to explain it to the user tha they have not passed correct arguments
// to the command line
var ErrWrongArgument error = fmt.Errorf("You may have used wrong argument/flag. Please use command `accurics help` to get a better idea")

// ErrAPICall -> User defined error to explain it to the user that something may have gon wrong with
// the server and that they should try again later
var ErrAPICall error = fmt.Errorf("Something went wrong with the server. Please try again, later")
