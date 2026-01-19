package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type task struct {
	Desc	string
	Status	int 
	Created time.Time
	Updated time.Time
	Next	*task
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
		for p.Next != nil {
			p = p.Next
		}
	} else {
		// scroll amnt times
		for amnt > 0 {
			p = p.Next
		}
	}
}

func main() {
	var args []string = os.Args
	amnt := len(args)
	if (amnt < 2) {
		// print help
		return
	}
	
	// decode 
	var head task
	var bytes_read int
	buffer := make([]byte, 1024)

	f, err := os.Open("tasks.JSON")
	if err == nil {
		bytes_read, err = f.Read(buffer)
		if err != nil {
			fmt.Println("read error: ", err)
			return
		}
		err = json.Unmarshal(buffer[:bytes_read], &head)
		if err != nil {
			fmt.Println("unmarshal error: ", err)
			return
		}
	}
	t := &head
	
	switch args[1] {
		case "add":
			if check(amnt, 3) == false {
				return
			}

			next := &task {
				Desc: args[2],
				Status: 0,
				Created: time.Now(),
				Updated: time.Now(),
				Next: nil,
			}
			t.Next = next

			fmt.Println("added", args[2])
			

		case "update":
			if check(amnt, 4) == false {
				return
			}

			id, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error: ", err)
				return
			}
			scroll(t, id)
			t.Updated = time.Now()	// update time


		case "delete":
			if check(amnt, 3) == false {
				return
			}

			id, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error: ", err)
				return
			}
			scroll(t, id-1)
			t.Next = t.Next.Next	// remove pointer


		case "mark":
			if check(amnt, 4) == false {
				return
			}

			id, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error: ", err)
				return
			}

			scroll(t, id)
			// mark status


		case "list":
			if check(amnt, 2) == false {
				return
			}

			p := &head;
			for p != nil {
				// print properties
				p = p.Next
			}


		default:
			// print help
			fmt.Println("help message")
			return

		// encode 
		out, oErr := os.Create("tasks.JSON")
		if oErr != nil {
			fmt.Println("create error: ", oErr)
			return
		}

		for t != nil {
			// somehow write each separate struct in the linked list to the same file
			t = t.Next
		}

		_, err = out.Write(buf)
		if err != nil {
			fmt.Println("write error: ", err)
			return
		}
	}
}

