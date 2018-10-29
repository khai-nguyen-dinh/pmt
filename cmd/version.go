package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func VersionPmt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of pmt.",
		Long:  `All software has versions. This is PMT's.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("PMT version v0.1 -- HEAD")
		},
	}
	return cmd
}
