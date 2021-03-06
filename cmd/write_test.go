package cmd

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	tmpFolder      = ".." + string(os.PathSeparator) + ".tmp"
	testDataFolder = ".." + string(os.PathSeparator) + "test-data"
)

func TestWriteFileNotFound(t *testing.T) {
	rootCmd.SetArgs([]string{"write", "foo", "bar", "--file", "does-not-exist.properties"})
	err := rootCmd.Execute()
	ce, ok := err.(*CmdError)
	if !ok {
		t.Errorf("No command error returned %s", err)
	}
	if ok && ce != &errCannotOpenFile {
		t.Errorf("Unexpected return of %s", err)
	}
}

func TestSimpleWrite(t *testing.T) {
	fileName := "basic.properties"
	err := createTemporaryFile(fileName, testDataFolder, tmpFolder)
	if err != nil {
		t.Errorf("Unexpected error when setting up test %s", err)
	}

	rootCmd.SetArgs([]string{"write", "foo", "baz", "--file", tmpFolder + string(os.PathSeparator) + fileName})
	err = rootCmd.Execute()
	if err != nil {
		t.Errorf("Unexpected error when writing %s", err)
	}
}

func createTemporaryFile(inputFile string, inputFolder string, tmpFolder string) error {
	os.RemoveAll(tmpFolder)
	err := os.Mkdir(tmpFolder, 0755)
	if err != nil {
		return err
	}

	input, err := ioutil.ReadFile(inputFolder + string(os.PathSeparator) + inputFile)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(tmpFolder+string(os.PathSeparator)+inputFile, input, 0644)
	if err != nil {
		return err
	}

	return nil
}
