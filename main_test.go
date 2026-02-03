package main

import (
	"os"
	"testing"
)

func TestOpenEmpty(t *Testing.T) {
	_ := os.Remove("tasks.JSON")
	f, fileExists := os.Open("tasks.JSON")
}

func TestReadEmpty(t *Testing.T) {
	_ := os.Remove("tasks.JSON")
	//
}

func TestWriteEmpty(t *Testing.T) {
	_ := os.Remove("tasks.JSON")
	//
}
