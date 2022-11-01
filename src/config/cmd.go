package config

import (
	"github.com/adiet95/moonlay-academy-technical-test/src/database"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "Simple backend login & register",
}

func init() {
	initCommand.AddCommand(ServeCmd)
	initCommand.AddCommand(database.MigrateCmd)

}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
