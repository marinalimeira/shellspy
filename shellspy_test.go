package shellspy_test

import (
	"bytes"
	"github.com/marinalimeira/shellspy"
	"strings"
	"testing"
)

func TestReads(t *testing.T) {
	reader := strings.NewReader("echo hi!")

	out := bytes.Buffer{}
	shellspy.Reads(reader, &out)

	got := out.String()
	want := "hi!"

	if got != want {
		t.Fatalf("%s and %s are not equal", got, want)
	}
}
