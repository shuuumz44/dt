package main

import (
	"os"
	"fmt"
	"time"
	"encoding/json"
	"strings"
)

type task struct {
	id int
	desc string
	status string
	created time.Time
	updated time.Time
}

func main() {
	var args []string = os.Args
	amnt := len(args)
	if amnt != 2 || args != 3 {
		fmt.Println("invalid number of arguments")
		return
	}
	fmt.Println(args[1])

	switch args[1] {
		case "add":
			//
		case "update":
			//
		case "delete":
			//
		case "mark":
			//
		case "done":
			//
		case "list":
			//
	}
}

