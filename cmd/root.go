package cmd

import (
	"os"

	"github.com/magiconair/properties"
	"github.com/spf13/cobra"
)

var (
	file      string
	useLatin1 bool = false
	rootCmd        = &cobra.Command{
		Use:   "props command",
		Short: "Reads or writes properties from and to properties files",
		Long:  ``,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&file, "file", "", "Properties file. If not set, STDIN/STDOUT is used")
	rootCmd.PersistentFlags().BoolVar(&useLatin1, "latin1", false, "Select Latin-1 encoding instead of the default UTF-8")
	rootCmd.AddCommand(writeCmd)
	rootCmd.AddCommand(readCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ce, ok := err.(*CmdError)
		if ok {
			os.Exit(ce.ReturnValue)
		}
		os.Exit(1)
	}
}

func getEncoding() properties.Encoding {
	if useLatin1 {
		return properties.ISO_8859_1
	}
	return properties.UTF8
}
