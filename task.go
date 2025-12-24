package main

import (
	"os"
	"fmt"
	"time"
	// "encoding/json"
)

type task struct {
	id int
	desc string
	status int 
	created time.Time
	updated time.Time
	next *task
}

func main() {
	var args []string = os.Args
	amnt := len(args)
	if amnt != 3 {
		fmt.Println("invalid number of arguments")
		return
	}
	
	var head task
	t := &head
	
	key := 1
	switch args[1] {
		case "add":
			fmt.Println("added", args[2])
			
			for t.next != nil {
				t = t.next
			}
			next := task {
				id: key,
				desc: args[2],
				status: 0,
				created: time.Now(),
				updated: time.Now(),
				next: nil,
			}
			t.next = &next
			key++
			
		case "update":
			//
		case "delete":
			//
		case "mark":
			//
		case "done":
			//
		case "list":
			p := &head;
			for p != nil {
				// print properties
				p = p.next;
			}
	}
}

