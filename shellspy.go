package shellspy

import (
	"bufio"
	"fmt"
	"io"
)

// Your implementation goes here!

func Reads(in io.Reader, out io.Writer) {
	input := bufio.NewScanner(in) // we know its a scanner, but it reads an input

	// 1. execute the command and output it.
	for input.Scan() {
		fmt.Fprint(out, input.Text())
	}
}
