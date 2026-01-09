package main

import (
	"encoding/json"
	"io"
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
	f, err := os.Open("tasks.JSON")
	if err == nil {
		dec := json.NewDecoder(f)
		for {
			if err := dec.Decode(head); err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("decode error")
				return
			}
			// check JSON for now
			fmt.Println(head)
		}

	}
	t := &head
	
	switch args[1] {
		case "add":
			if check(amnt, 3) == false {
				return
			}

			next := task {
				Desc: args[2],
				Status: 0,
				Created: time.Now(),
				Updated: time.Now(),
				Next: nil,
			}
			t.Next = &next

			fmt.Println("added", args[2])
			
		case "update":
			if check(amnt, 4) == false {
				return
			}

			id, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error:", err)
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
				fmt.Println("string convert error:", err)
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
				fmt.Println("string convert error:", err)
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

		// encode file back into JSON
		out, err := os.Create("tasks.JSON")
		if err != nil {
			fmt.Println("file could not be written to.") 
		}

		enc := json.NewEncoder(out)
		if err := enc.Encode(t); err != nil {
			fmt.Println("writing error")
		}
	}
}

