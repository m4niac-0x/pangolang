// unix command execution requiring user input

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	executeUnixCmd("ls", "-l")

}

func executeUnixCmd(c string, a ...string) {
	cmd := exec.Command(c, a...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		checkError(err)
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		panic(err)
	}
}
