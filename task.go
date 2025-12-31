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

func check(needed int, amnt int) bool {
	if (needed==amnt) {
		return true
	}

	fmt.Println("invalid number of arguments")
	return false
}

func scroll(p *task, amnt int) {
	// scrolls pointer p to specified value. (args[2] or end)
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
	if (amnt <= 1) {
		// print help
		return
	}
	
	var head task
	t := &head
	
	switch args[1] {
		case "add":
			rc := check(amnt, 3)
			if (rc == false) {
				return
			}

			scroll(t, 0)
			next := task {
				desc: args[2],
				status: 0,
				created: time.Now(),
				updated: time.Now(),
				next: nil,
			}
			t.next = &next

			fmt.Println("added", args[2])
			
		case "update":
			rc := check(amnt, 4)
			if (rc == false) {
				return
			}

			id, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error:", err)
				return
			}
			scroll(t, id)
			t.updated = time.Now()	// update time

		case "delete":
			rc := check(amnt, 3)
			if (rc == false) {
				return
			}
			id, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error:", err)
				return
			}
			scroll(t, id-1)
			t.next = t.next.next	// remove pointer

		case "mark":
			rc := check(amnt, 4)
			if (rc == false) {
				return
			}

			id, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error:", err)
				return
			}

			scroll(t, id)
			// mark status

		case "list":
			rc := check(amnt, 2)
			if (rc == false) {
				return
			}

			p := &head;
			for p != nil {
				// print properties
				p = p.next
			}

		default:
			// print help
			fmt.Println("help message")
	}
}

