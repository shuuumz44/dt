package main

import (
	"os"
	"fmt"
	"time"
	"strconv"
	// "encoding/json"
)

type task struct {
	desc string
	status int 
	created time.Time
	updated time.Time
	next *task
}

func scroll(p *task, amnt int) {
	// scrolls t pointer to specified value. (args[2] or end)
	if amnt==0 {
		// treat '0' as scroll to end
		for p.next != nil {
			p = p.next
		}
	} else {
		// scroll amnt times
		for amnt > 0 {
			p = p.next
		}
	}
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
	
	switch args[1] {
		case "add":
			fmt.Println("added", args[2])
			
			scroll(t, 0)
			next := task {
				desc: args[2],
				status: 0,
				created: time.Now(),
				updated: time.Now(),
				next: nil,
			}
			t.next = &next
			
		case "update":
			id, _ := strconv.Atoi(args[2])
			scroll(t, id)
			t.updated = time.Now()

		case "delete":
			id, _ := strconv.Atoi(args[2])
			scroll(t, id)
			// remove pointer

		case "mark":
			id, _ := strconv.Atoi(args[2])
			scroll(t, id)
			// mark status

		case "list":
			p := &head;
			for p != nil {
				// print properties
				p = p.next;
			}
	}
}

