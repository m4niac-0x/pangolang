// unix command execution requiring user input

package main

import (
	"os"
	"os/exec"
)

func main() {

	// To run any system commands. EX: Cloud Foundry CLI commands: `CF login`
	cmd := exec.Command("cmd", "arg1")

	// Sets standard output to cmd.stdout writer
	cmd.Stdout = os.Stdout

	// Sets standard input to cmd.stdin reader
	cmd.Stdin = os.Stdin

	cmd.Run()

}
