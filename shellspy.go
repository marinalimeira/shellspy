package shellspy

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

// Object in Go means the struct.

func ExecCommand(command *exec.Cmd) (string, error) {
	output, err := command.CombinedOutput()

	return string(output), err
}

// CommandFromString returns a CommandObject from input line.
func CommandFromString(command string) *exec.Cmd {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	return cmd
}

func Reads(in io.Reader, out io.Writer) {
	input := bufio.NewScanner(in) // we know its a scanner, but it reads an input

	fmt.Fprint(out, "$ ")

	// 1. execute the command and output it.
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}

		command := CommandFromString(input.Text())
		output, err := ExecCommand(command)
		if err != nil {
			fmt.Fprintln(out, err)
		}
		// out < `hi!`
		fmt.Fprint(out, output)

		fmt.Fprint(out, "$ ")
	}
}
