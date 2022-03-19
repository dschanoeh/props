package cmd

type CmdError struct {
	Message     string
	ReturnValue int
}

var errCannotOpenFile = CmdError{Message: "cannot open the properties file", ReturnValue: 2}
var errPropertyNotFound = CmdError{Message: "the property wasn't found", ReturnValue: 3}
var errCouldNotSetProperty = CmdError{Message: "could not set the property", ReturnValue: 4}
var errCouldNotOpenForWrite = CmdError{Message: "could not open properties file for writing", ReturnValue: 5}
var errCannotReadStdin = CmdError{Message: "could not read from stdin", ReturnValue: 6}

func (m *CmdError) Error() string {
	return m.Message
}
