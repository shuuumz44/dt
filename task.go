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
}

func check(needed int, amnt int) bool {
	if (needed==amnt) {
		return true
	}

	fmt.Println("invalid number of arguments")
	return false
}

func listTask(t *task) {
	fmt.Println(t.Desc)
	fmt.Println("status:		", t.Status)
	fmt.Println("began:		", t.Created)
	fmt.Println("last updated:	", t.Updated)
	fmt.Println()
}

func main() {
	// macros
	ARR_MAX := 1000
	help := `
	dt: task manager CLI
	usage: 
		dt [add] [update] [delete] [mark] [list]
	examples:
		add a new task called "get milk":
		dt add	"get milk"

		update the status of task "exercise":
		dt update exercise

		mark task "meditate" as in progress:
		dt mark meditate 1

		list all existing tasks:
		dt list [todo | doing | done]
		*add 1 of the arguments in brackets to filter the list by status.

	`

	var args []string = os.Args
	amnt := len(args)
	if (amnt < 2) {
		fmt.Println(help)
		return
	}
	
	// decode 
	arr := make([]task, 0, ARR_MAX)
	buffer := make([]byte, 1024)
	numTasks := 0

	f, err := os.Open("tasks.JSON")
	if err == nil {
		bytes_read, err := f.Read(buffer)
		if err != nil {
			fmt.Println("read error: ", err)
			return
		}

		if json.Valid(buffer[:bytes_read]) == false {
			fmt.Println("invalid JSON");
			return
		}

		err = json.Unmarshal(buffer[:bytes_read], &arr)
		if err != nil {
			fmt.Println("unmarshal error: ", err)
			return
		}
		numTasks = len(arr)
	}
	
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
			}

			// append to arr
			arr = append(arr, next)
			fmt.Println("added:", args[2])
			

		case "update":
			if check(amnt, 4) == false {
				return
			}

			name := args[3]
			id, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error: ", err)
				return
			}

			arr[id].Desc = name;
			arr[id].Updated = time.Now()


		case "delete":
			if check(amnt, 3) == false {
				return
			}

			id, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error: ", err)
				return
			}

			next := id+1
			arr = append(arr[:id], arr[next:numTasks]...)

		case "mark":
			if check(amnt, 4) == false {
				return
			}

			newStatus, err := strconv.Atoi(args[3])
			id, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error: ", err)
				return
			}

			arr[id].Status = newStatus


		case "list":
			if check(amnt, 2) == false {
				return
			}

			for i:=0; i < numTasks; i++ {
				listTask(&arr[i])
			}

		default:
			// print help
			fmt.Println(help)
			return
	}


	// encode 
	out, err := os.Create("tasks.JSON")
	if err != nil {
		fmt.Println("create error: ", err)
		return
	}

	// marshal
	buffer, err = json.Marshal(&arr)
	if err != nil {
		fmt.Println("marshal error: ", err)
		return
	}

	// write to JSON
	_, err = out.Write(buffer)
	if err != nil {
		fmt.Println("write error: ", err)
		return
	}
}

