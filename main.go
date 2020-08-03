package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		helpCmd()
		return
	}

	switch os.Args[1] {
	case HELP1, HELP2, HELP3, HELP4:
		helpCmd()
	case LISTCMD:
		err := listCmd()
		handle(err)
	case CREATECMD:
		if len(os.Args) < 8 {
			helpCmd()
			return
		}
		err := createCmd(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6], os.Args[7])
		handle(err)
	case DELETECMD:
		if len(os.Args) < 4 {
			helpCmd()
			return
		}
		err := delCmd(os.Args[2], os.Args[3])
		handle(err)
	case SHOWCMD:
		if len(os.Args) < 4 {
			helpCmd()
			return
		}
		err := showCmd(os.Args[2], os.Args[3])
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
