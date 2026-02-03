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

func ListTasks(t []task, amnt, status int) {
	skipped := 0
	for i := 0; i < amnt; i++ {
		var marked string
		switch t[i].Status {
			case 0:
				marked = "todo"
			case 1:
				marked = "doing"
			default:
				marked = "done"
		}
		
		if status < 3 && t[i].Status != status {
			skipped++
			continue
		}

		fmt.Println(t[i].Desc)
		fmt.Println("task ", i+1)
		fmt.Println("status:		", marked)
		fmt.Println("began:		", t[i].Created.Format("Mon Jan 02 3:04PM"))
		fmt.Println("last updated:	", t[i].Updated.Format("Mon Jan 02 3:04PM"))
		fmt.Println()
	}
	if skipped==amnt {
		fmt.Println("no tasks found.")
	}
}

func Add(arr *[]task, name string) {
	next := task {
		Desc: name,
		Status: 0,
		Created: time.Now(),
		Updated: time.Now(),
	}

	fmt.Println("added:", name)
	*arr = append(*arr, next)
}

func Update(arr *[]task, num string, name string) error {
	id, err := strconv.Atoi(num)
	if (err != nil) {
		fmt.Println("string convert error: ", err)
		return err
	}

	(*arr)[id-1].Desc = name;
	(*arr)[id-1].Updated = time.Now()

	return nil
}

func Delete(arr *[]task, num string, tasks int) error {
	id, err := strconv.Atoi(num)
	if (err != nil) {
		fmt.Println("string convert error: ", err)
		return err
	}
	
	next := id
	*arr = append((*arr)[:id-1], (*arr)[next:tasks]...)

	return nil
}

func Mark(arr *[]task, num, current string) error {
	newName, err1 := strconv.Atoi(current)
	if (err1 != nil ) {
		fmt.Println("string convert error: ", err1)
		return err1 
	}
	id, err2 := strconv.Atoi(num)
	if (err2 != nil ) {
		fmt.Println("string convert error: ", err2)
		return err2
	}

	(*arr)[id-1].Status = newName

	return nil
}

func main() {
	// macros
	ARR_MAX := 1000
	help :=
`dt: task manager CLI
usage: 
	dt [add] [update] [delete] [mark] [list]
examples:
	add a new task called "get milk":
	dt add	"get milk"

	update the status of task "exercise":
	dt update exercise

	mark task "meditate" as in progress:
	dt mark [task id] [status]
	* status is an integer between 0 and 2.
	these represent states "todo", "doing", and "done".

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

	f, fileExists := os.Open("tasks.JSON")
	if fileExists == nil {
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
			if amnt != 3 {
				fmt.Println("invalid argument count")
				return 
			}

			Add(&arr, args[2]);

		case "update":
			if amnt != 4 {
				fmt.Println("invalid argument count")
				return 
			}
			if fileExists != nil {
				fmt.Println("no tasks to update.")
				return 
			}

			err := Update(&arr, args[2], args[3]) 

			if err != nil {
				fmt.Println("failed to update.")
				return
			}


		case "delete":
			if amnt != 3 {
				fmt.Println("invalid argument count")
				return 
			}
			if fileExists != nil {
				fmt.Println("no tasks to delete.")
				return 
			}

			err := Delete(&arr, args[2], numTasks)

			if (err != nil) {
				fmt.Println("could not delete task.")
				return 
			}


		case "mark":
			if amnt != 4 {
				fmt.Println("invalid argument count")
				return 
			}
			if fileExists != nil {
				fmt.Println("no tasks to mark.")
				return 
			}

			err := Mark(&arr, args[2], args[3])
			if err != nil {
				fmt.Println("could not mark task.")
				return
			}

		case "list":
			if amnt == 2 {
				ListTasks(arr[:numTasks], numTasks, 3)
			} else if amnt == 3 {

				status, err := strconv.Atoi(args[2])
				if err != nil {
					fmt.Println("string convert error: ", err)
					return 
				}
				ListTasks(arr[:numTasks], numTasks, status)

			} else {
				fmt.Println("invalid argument count")
				return 
			}


		default:
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

