package main

import (
    "encoding/json"
	"fmt"
    "os"
)

type character struct {
	Name 		string
	Limbs 		int
	Occupation 	string
	Eyes 		string
	Temperament string
}

func main() {
	// macros
	ARR_MAX := 1000

	f, err := os.Open("character.JSON")
	if err != nil {
		fmt.Println("read error: ", err)
	}

	// read file into buffer
	buf := make([]byte, 1024)
	bytes_read, err := f.Read(buf);
	if err != nil {
		fmt.Println("read error: ", err)
		return
	}

	if json.Valid(buf[:bytes_read]) == false {
		fmt.Println("invalid JSON.")
		return
	}

	// make file to write to
	out, cErr := os.Create("character.JSON")
	if cErr != nil {
		fmt.Println("create error: ", cErr)
		return
	}

	// read JSON into struct array
	arr := make([]character, 0, ARR_MAX)
	if err = json.Unmarshal(buf[:bytes_read], &arr); err != nil {
		fmt.Println("unmarshal error: ", err)
	}
	characterAmnt := len(arr)

	// manipulate struct as desired
	fmt.Println("amount of characters: ", characterAmnt)
	arr[1].Name = "Patrick"

	// marshal
	buf, err = json.Marshal(&arr)
	if err != nil {
		fmt.Println("marshal error: ", err)
		return
	}

	// write to file from buffer
	_, err = out.Write(buf)
	if err != nil {
		fmt.Println("write error: ", err);
		return
	}
}
