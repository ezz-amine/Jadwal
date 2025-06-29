package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/ezz-amine/Jadwal/pkg/application"
	"github.com/ezz-amine/Jadwal/pkg/core"

	"github.com/spf13/cobra"
)

type CreateTableCommand struct {
	application.Handler
}

func (c CreateTableCommand) Args(app application.Application, cmd *cobra.Command, args []string) error {
	var err error = nil

	if err = cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), core.TIMEOUT)
	defer cancel()

	title := strings.Join(args, " ")
	_, err = app.Queries.GetTableByTitle(ctx, title)
	if err == nil {
		return fmt.Errorf("table titled '%s' already exists", title)
	}

	return nil
}

func (c CreateTableCommand) Run(app application.Application, cmd *cobra.Command, args []string) error {
	title := strings.Join(args, " ")
	fmt.Println(title)

	ctx, cancel := context.WithTimeout(context.Background(), core.TIMEOUT)
	defer cancel()

	_, err := app.Queries.CreateTable(ctx, title)
	if err != nil {
		return fmt.Errorf("error occurred when creating table '%s': %w", title, err)
	}

	fmt.Printf("table titled '%s' created", title)
	return nil
}

var createTableCmd = application.NewCommand(
	CreateTableCommand{},
	application.CommandDefinition{
		Use:   "create",
		Short: "create a new table:'Jadwal' for your todos",
	},
)

func init() {
	tableCmd.AddCommand(createTableCmd)
}
