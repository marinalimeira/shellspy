package shellspy_test

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"github.com/marinalimeira/shellspy"
	"path/filepath"
	"strings"
	"testing"
)

func TestReads_ExecutesTheInputsAsCommands(t *testing.T) {
	reader := strings.NewReader("echo hi!")

	out := bytes.Buffer{}
	shellspy.Reads(reader, &out)

	got := out.String()
	want := "hi!\n"

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

func TestReads_FinishesOnExit(t *testing.T) {
	// TODO

}

func TestCommandFromString_ReturnsCommandObjectFromInputLine(t *testing.T) {
	// Given a string,
	// It will return the command to be executed, and it's args.

	testCase := []struct {
		input    string
		want     []string
		basePath string
	}{
		{
			input:    "echo hi!",
			want:     []string{"echo", "hi!"},
			basePath: "echo",
		},
		{
			input:    "echo",
			want:     []string{"echo"},
			basePath: "echo",
		},
		{
			input:    "abc",
			want:     []string{"abc"},
			basePath: "abc",
		},
		{
			input:    "",
			want:     []string{""},
			basePath: ".",
		},
	}

	for _, test := range testCase {
		got := shellspy.CommandFromString(test.input)

		// In this case, no need to compare the whole struct, only the attrs that are changing

		if !cmp.Equal(test.want, got.Args) {
			t.Errorf("Input: %q\nDiff: %s", test.input, cmp.Diff(test.want, got.Args))
		}

		if filepath.Base(got.Path) != test.basePath {
			t.Errorf("Path should end with %q. Got: %q\nFull path: %q", test.basePath, filepath.Base(got.Path), got.Path)
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
	//command := exec.Command("echo", "hi!")
	//
	input := "echo hi!"
	command := shellspy.CommandFromString(input)

	expectedBasePath := "echo"
	if filepath.Base(command.Path) != expectedBasePath {
		t.Errorf("Path should end with %q. Got full path %q", expectedBasePath, command.Path)
	}

	// 3. Start from middle
	got, err := shellspy.ExecCommand(command)
	if err != nil {
		t.Error(err)
	}

	want := "hi!\n"

	if got != want {
		t.Errorf("Input: %q\n Diff: %s", input, cmp.Diff(want, got))
	}
}
