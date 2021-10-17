package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := recordUserInput("Do you want to know the Answer to the Ultimate Question of Life, The Universe, and Everything ? (y/[n]) ")
	switch r {
	case "y":
		fmt.Println("Use Arch BTW")
	case "null":
		fmt.Println("Avoiding...")
	case "n":
		fmt.Println("Avoiding...")
	default:
		main()
	}
}
func recordUserInput(r string) string {
	request := r
	var a string
	fmt.Printf("%s ", request)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		a := scanner.Text()
		// fmt.Println(r)
		if len(a) < 1 {
			a = "null"
			return a
		} else {
			return a
		}
	}
	return a
}
