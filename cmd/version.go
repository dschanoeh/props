package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version = ""
	Commit  = ""
	Date    = ""
	BuiltBy = ""
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "props %s, commit %s, built at %s by %s\n", Version, Commit, Date, BuiltBy)
	},
}
