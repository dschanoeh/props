package cmd

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(licenseCmd)
}

var (
	licenseCmd = &cobra.Command{
		Use:   "license",
		Short: "Print license information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "props license information:\n\n%s\n", propsLicense)
			fmt.Fprintf(cmd.OutOrStdout(), "================================================================================\n")
			fmt.Fprintf(cmd.OutOrStdout(), "Third party license information:\n\n%s\n", thirdPartyLicenses)
		},
	}

	//go:embed licenses/3rd-party-licenses.txt
	thirdPartyLicenses string

	//go:embed licenses/LICENSE
	propsLicense string
)
