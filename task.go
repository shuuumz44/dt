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

func main() {
	// macros
	ARR_MAX := 1000

	var args []string = os.Args
	amnt := len(args)
	if (amnt < 2) {
		// print help
		return
	}
	
	// decode 
	arr := make([]task, 0, ARR_MAX)
	buffer := make([]byte, 1024)
	// numTasks := 0

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
		// numTasks = len(arr)
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

			_, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error: ", err)
				return
			}
			// arr[id].Updated = time.Now()	// update time


		case "delete":
			if check(amnt, 3) == false {
				return
			}

			_, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error: ", err)
				return
			}
			// this one will take a bit more effort

		case "mark":
			if check(amnt, 4) == false {
				return
			}

			_, err := strconv.Atoi(args[2])
			if (err != nil) {
				fmt.Println("string convert error: ", err)
				return
			}

			// mark status
			// arr[i].Status = args[3]


		case "list":
			if check(amnt, 2) == false {
				return
			}

			// i believe you can use a slice for this.
			// for arr[:len] {print data}

		default:
			// print help
			fmt.Println("help message")
			return
	}

	fmt.Println("first task: ", arr[0].Desc)

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

