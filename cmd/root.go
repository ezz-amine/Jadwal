package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/ezz-amine/Jadwal/pkg/core"
	"github.com/ezz-amine/Jadwal/pkg/output"
	"github.com/ezz-amine/Jadwal/pkg/sqlc"
	"github.com/spf13/cobra"
)

// remove those use a custom command stucture
var (
	selectTable *sqlc.TodoTable = nil
	queries     *sqlc.Queries   = nil
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Jadwal",
	Short: "A simple yet powerful task scheduler built in Go, organizing your todos in a clean, table-like system. Plan, track, and conquer your day with structured efficiency.",
	Long: `A Go experiment in structured task management.
Jadwal organizes todos across multiple tables or lists—like a lightweight database for your goals. Designed primarily as a learning project (not a polished tool), it’s a playground for exploring Go, architecture, and open-source collaboration.`,
	Args: func(cmd *cobra.Command, args []string) error {
		var (
			err   error  = nil
			title string = strings.Join(args, " ")
		)

		ctx, cancel := context.WithTimeout(context.Background(), core.TIMEOUT)
		defer cancel()

		queries, err = core.GetQueries()
		if err != nil {
			return fmt.Errorf("database err: %w", err)
		}

		if err = cobra.MinimumNArgs(1)(cmd, args); err == nil {
			table, err := queries.GetTableByTitle(ctx, title)
			selectTable = &table
			if err != nil {
				return fmt.Errorf("database err: %w", err)
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		termContext := core.NewTermContext()
		if selectTable == nil {
			output.DisplayAllTables(termContext, queries)
		} else {
			output.DisplayTable(termContext, queries, *selectTable)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Jadwal.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
