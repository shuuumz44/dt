package main

import (
	"os"
	"fmt"
)

type task struct {
	id int
	desc string
	status string
	created string
	updated string
}

func main() {
	var args []string = os.Args
	amnt := len(args)
	if amnt != 2 {
		fmt.Println("invalid number of arguments")
		return
	}
	fmt.Println(args[1])
}

