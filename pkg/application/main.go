package application

import (
	"fmt"
	"os"

	"github.com/ezz-amine/Jadwal/pkg/core"
	"github.com/ezz-amine/Jadwal/pkg/sqlc"
	"github.com/spf13/cobra"
)

type Application struct {
	Queries     *sqlc.Queries
	TermContext *core.TermContext
}

type Handler interface {
	Args(app Application, cmd *cobra.Command, args []string) error
	Run(app Application, cmd *cobra.Command, args []string) error
}

type CommandDefinition struct {
	Use   string
	Short string
	Long  string
}

func NewCommand(handler Handler, definition CommandDefinition) *cobra.Command {
	termContext := core.NewTermContext()
	queries, err := core.GetQueries()
	if err != nil {
		ExitWithError(err)
	}
	app := Application{queries, termContext}
	return &cobra.Command{
		Use:   definition.Use,
		Short: definition.Short,
		Long:  definition.Long,
		Args: func(cmd *cobra.Command, args []string) error {
			return handler.Args(app, cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			err = handler.Run(app, cmd, args)
			if err != nil {
				ExitWithError(err)
			}
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
}

func ExitWithError(err error) {
	fmt.Fprintf(os.Stderr, "\033[31mError:\033[0m %v\n", err)
	os.Exit(1)
}
