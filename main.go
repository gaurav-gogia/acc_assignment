package main

import (
	"log"
	"os"
)

func main() {
	// Get DB Connection. Create employee table if it doesn't exist
	db, err := initdb()
	handle(err)
	defer db.Close()

	if len(os.Args) < 2 {
		helpCmd()
		os.Exit(0)
	}

	switch os.Args[1] {
	case HELP1, HELP2, HELP3, HELP4:
		helpCmd()
	case LISTCMD:
		err := listCmd(db)
		handle(err)
	case CREATECMD:
		err := createCmd(db, os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6], os.Args[7])
		handle(err)
	case DELETECMD:
		err := delCmd(db, os.Args[2], os.Args[3])
		handle(err)
	case SHOWCMD:
		err := showCmd(db, os.Args[2], os.Args[3])
		handle(err)
	default:
		helpCmd()
	}
}

func handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
