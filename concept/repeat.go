package main

import (
    "encoding/json"
	"fmt"
    "log"
    "os"
)

type character struct {
	Name 		string
	Limbs 		int
	Occupation 	string
	Eyes 		string
	Temperament string
}

/* This is the guideline for using the encoder and decoder methods
to manipulate JSON files, through structs. */
func main() {
	f, err := os.Open("character.JSON")
	if err != nil {
		fmt.Println("file could not be read from.")
	}

    dec := json.NewDecoder(f)

	out, err := os.Create("out.JSON")
	if err != nil {
		fmt.Println("file could not be written to.")
	}
    enc := json.NewEncoder(out)

	var v character
	for {
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }

        if err := enc.Encode(&v); err != nil {
            log.Println(err)
        }
    }
}
