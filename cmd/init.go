package cmd

import (
	"fmt"
	"willbehn/what-the-terminal/internal"

	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long:  `fiks senere`,
	RunE: func(cmd *cobra.Command, args []string) error {

		db, err := internal.OpenDB()

		if err != nil {
			return err
		}

		defer db.Close()

		if _, err := db.Exec(`PRAGMA journal_mode=WAL;`); err != nil {
			return err
		}

		if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS commands (
								id INTEGER PRIMARY KEY,
								ts INTEGER NOT NULL,
								shell TEXT ,
								dir TEXT ,
								repo TEXT,
								branch TEXT,
								cmd TEXT ,
								exit_code INTEGER ,
								duration_ms INTEGER );`); err != nil {
			return err
		}

		if _, err := db.Exec(`CREATE INDEX IF NOT EXISTS idx_commands_ts ON commands(ts DESC);`); err != nil {
			return err
		}

		fmt.Println("database for command history successuflly inited <3")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}
