package cmd

import (
	"bufio"
	"os"

	"github.com/magiconair/properties"
	"github.com/spf13/cobra"
)

var (
	writeCmd = &cobra.Command{
		Use:   "write property value",
		Short: "Write a single property to the properties file",
		Long: `Will modify the properties file in place. If the property didn't
				exist yet, it will be added at the end of the file. If it did
				exist, it will be overwritten.`,
		Args: cobra.ExactArgs(2),
		RunE: Write,
	}
	create bool
)

func Write(cmd *cobra.Command, args []string) error {

	propertyName := args[0]
	propertyValue := args[1]
	encoding := getEncoding()

	p, err := readProperties(cmd.InOrStdin())
	if err != nil {
		/* If the file cannot be opened and the create flag is set,
		 * use an empty properties object to start with.
		 */
		ce, ok := err.(*CmdError)
		if create && ok && ce == &errCannotOpenFile {
			p = properties.NewProperties()
		} else {
			return &errCannotOpenFile
		}
	}

	err = p.SetValue(propertyName, propertyValue)
	if err != nil {
		return &errCouldNotSetProperty
	}

	var f *os.File
	if file != "" {
		f, err = os.Create(file)
		if err != nil {
			return &errCouldNotOpenForWrite
		}
		defer f.Close()
	} else {
		f = os.Stdout
	}

	w := bufio.NewWriter(f)
	p.WriteComment(w, "#", encoding)
	w.Flush()

	return nil
}

func init() {
	writeCmd.Flags().BoolVar(&create, "create", false, "When set, the properties file will be created if it doesn't exist yet")
}
