package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestFileNotFound(t *testing.T) {
	rootCmd.SetArgs([]string{"read", "foo", "--file", "does-not-exist.properties"})
	err := rootCmd.Execute()
	ce, ok := err.(*CmdError)
	if !ok {
		t.Errorf("No command error returned %s", err)
	}
	if ok && ce != &errCannotOpenFile {
		t.Errorf("Unexpected return of %s", err)
	}
}

func TestReadProperty(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"read", "test", "--file", "../test-data/basic.properties"})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "test3" {
		t.Fatalf("expected \"%s\" got \"%s\"", "test3", string(out))
	}
}

func TestReadUTF8(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"read", "foo", "--file", "../test-data/utf8.properties"})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "🤓" {
		t.Fatalf("expected \"%s\" got \"%s\"", "🤓", string(out))
	}
}

func TestReadLatin1(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"read", "foo", "--file", "../test-data/utf8.properties", "--latin1"})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "ð¤" {
		t.Fatalf("expected \"%s\" got \"%s\"", "ð¤", string(out))
	}
}

func TestReadMultiline(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"read", "multiline", "--file", "../test-data/basic.properties"})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "foobar" {
		t.Fatalf("expected \"%s\" got \"%s\"", "foobar", string(out))
	}
}

func TestReadFromSTDIN(t *testing.T) {
	b := bytes.NewBufferString("")
	in := strings.NewReader("foo: bar")
	rootCmd.SetIn(in)
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"read", "foo"})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "bar" {
		t.Fatalf("expected \"%s\" got \"%s\"", "bar", string(out))
	}
}
