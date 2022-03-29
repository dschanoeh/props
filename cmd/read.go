package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/magiconair/properties"
	"github.com/spf13/cobra"
)

var (
	readCmd = &cobra.Command{
		Use:   "read property",
		Short: "Read a single property from the properties file",
		Args:  cobra.ExactArgs(1),
		RunE:  Read,
	}
)

func Read(cmd *cobra.Command, args []string) error {
	propertyName := args[0]

	p, err := readProperties()
	if err != nil {
		return err
	}

	propertyValue, found := p.Get(propertyName)

	if found {
		fmt.Fprint(cmd.OutOrStdout(), propertyValue)
		return nil
	} else {
		return &errPropertyNotFound
	}
}

func readProperties() (*properties.Properties, error) {
	var p *properties.Properties
	var err error
	encoding := getEncoding()

	if file != "" {
		p, err = properties.LoadFile(file, encoding)
		if err != nil {
			return nil, &errCannotOpenFile
		}
	} else {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, &errCannotReadStdin
		}
		p, err = properties.Load(bytes, encoding)
		if err != nil {
			return nil, &errCannotOpenFile
		}
	}

	return p, nil
}
