package shellspy

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

// Object in Go means the struct.

func ExecCommand(command *exec.Cmd) string {
	return ""
}

// CommandFromString returns a CommandObject from input line.
func CommandFromString(command string) exec.Cmd {
	args := strings.Split(command, " ")
	return exec.Cmd{Args: args}
}

func Reads(in io.Reader, out io.Writer) {
	input := bufio.NewScanner(in) // we know its a scanner, but it reads an input

	// 1. execute the command and output it.
	for input.Scan() {
		fmt.Fprint(out, input.Text())

	}
}
