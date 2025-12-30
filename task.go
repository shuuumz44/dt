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
	
	var head task
	t := &head
	
	switch args[1] {
		if amnt != 3 {
			fmt.Println("invalid number of arguments")
			return
		}

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
			if amnt != 4 {
				fmt.Println("invalid number of arguments")
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
			if amnt != 3 {
				fmt.Println("invalid number of arguments")
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
			if amnt != 4 {
				fmt.Println("invalid number of arguments")
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
			if amnt != 2 {
				fmt.Println("invalid number of arguments")
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

