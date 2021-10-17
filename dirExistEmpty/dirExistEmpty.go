package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	dir := "/path/to/dir"

	if !dirExist(dir) {
		fmt.Println("Directory doesn't exist")
	}
	e, err := dirIsEmpty(dir)
	checkError(err)
	if !e {
		fmt.Println("Directory insn't empty")
	}
}

func dirExist(d string) bool {
	_, err := os.Stat(d)
	var b bool
	if err == nil {
		b = true
	} else if os.IsNotExist(err) {
		b = false
	} else {
		checkError(err)
	}
	return b
}

func dirIsEmpty(d string) (bool, error) {
	f, err := os.Open(d)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		panic(err)
	}
}
