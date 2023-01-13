package shellspy_test

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"github.com/marinalimeira/shellspy"
	"os/exec"
	"strings"
	"testing"
)

func TestReads_ExecutesTheInputsAsCommands(t *testing.T) {
	reader := strings.NewReader("echo hi!")

	out := bytes.Buffer{}
	shellspy.Reads(reader, &out)

	got := out.String()
	want := "hi!"

	if got != want {
		// With %q it shows the string with quotes
		t.Fatalf("Want: %q\nGot: %q", want, got)
	}
}

func TestReads_IgnoresWhenItsEmpty(t *testing.T) {
	reader := strings.NewReader("")

	out := bytes.Buffer{}
	shellspy.Reads(reader, &out)

	got := out.String()
	want := ""

	if got != want {
		// With %q it shows the string with quotes
		t.Fatalf("Want: %q\nGot: %q", want, got)
	}
}

func TestCommandFromString_ReturnsCommandObjectFromInputLine(t *testing.T) {
	// Given a string,
	// It will return the command to be executed, and it's args.

	testCase := []struct {
		input string
		want  []string
	}{
		{
			input: "echo hi!",
			want:  []string{"echo", "hi!"},
		},
		{
			input: "echo",
			want:  []string{"echo"},
		},
		{
			input: "abc",
			want:  []string{"abc"},
		},
	}

	for _, test := range testCase {
		got := shellspy.CommandFromString(test.input)

		// In this case, no need to compare the whole struct, only the attrs that are changing

		if !cmp.Equal(test.want, got.Args) {
			t.Errorf("Input: %q\nDiff: %s", test.input, cmp.Diff(test.want, got.Args))
		}
	}
}

// 1. Write phrase with what will happen
func TestExecCommand_RunsAndReturnsTheOutputGivenOfCommand(t *testing.T) {
	// 2. Write test that does whatever "phrase" above says

	// it returns a pointer, something changes.
	// Entity.
	// Read-only.
	//
	command := exec.Command("echo", "hi!")

	// 3. Start from middle
	got := shellspy.ExecCommand(command)

	want := "hi!"

	if got != want {
		t.Errorf("Input: \"echo hi!\"\n Diff: %s", cmp.Diff(want, got))
	}
}
