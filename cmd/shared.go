package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func ExitWithError(cmd *cobra.Command, err error) {
	cmd.PrintErrf("\033[31mError:\033[0m %v\n", err)
	os.Exit(1)
}
