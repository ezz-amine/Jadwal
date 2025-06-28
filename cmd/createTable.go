/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/ezz-amine/Jadwal/pkg/core"
	// "github.com/ezz-amine/Jadwal/pkg/sqlc"
	"github.com/spf13/cobra"
)

// queries *sqlc.Queries = nil
var title *string = nil

// createTableCmd represents the createTable command
var createTableCmd = &cobra.Command{
	Use:   "createTable",
	Short: "create a new 'Jadwal' for your todos",
	Args: func(cmd *cobra.Command, args []string) error {
		var err error = nil

		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), core.TIMEOUT)
		defer cancel()

		if queries == nil {
			queries, err = core.GetQueries()
			if err != nil {
				return fmt.Errorf("database err: %w", err)
			}
		} else {
			fmt.Println("lreadi\n")
		}

		tempArgs := strings.Join(args, " ")
		title = &tempArgs
		_, err = queries.GetTableByTitle(ctx, *title)
		if err == nil {
			return fmt.Errorf("table with title '%s' already exists", *title)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(*title)
		queries.
	},
}

func init() {
	rootCmd.AddCommand(createTableCmd)
}
