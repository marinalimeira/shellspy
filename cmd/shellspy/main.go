package main

import (
	"github.com/marinalimeira/shellspy"
	"os"
)

// Your CLI goes here!

func main() {
	shellspy.Reads(os.Stdin, os.Stdout)
}
