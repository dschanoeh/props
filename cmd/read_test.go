package cmd

import (
	"testing"
)

func TestFileNotFound(t *testing.T) {
	file = "does-not-exist.properties"
	err := Read(readCmd, []string{"foo"})
	ce, ok := err.(*CmdError)
	if !ok {
		t.Errorf("No command error returned %s", err)
	}
	if ok && ce != &errCannotOpenFile {
		t.Errorf("Unexpected return of %s", err)
	}
}

func TestReadProperty(t *testing.T) {
	file = "../test-data/basic.properties"
	err := Read(readCmd, []string{"test"})
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func ExampleRead() {
	file = "../test-data/basic.properties"
	Read(readCmd, []string{"test"})
	// Output: test3
}

func ExampleRead_utf8() {
	file = "../test-data/utf8.properties"
	useUTF8 = true
	Read(readCmd, []string{"foo"})
	// Output: 🤓
}

func ExampleRead_multiline() {
	file = "../test-data/basic.properties"
	Read(readCmd, []string{"multiline"})
	// Output: foobar
}
