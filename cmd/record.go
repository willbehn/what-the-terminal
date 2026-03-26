package cmd

import (
	"willbehn/what-the-terminal/internal"
	"willbehn/what-the-terminal/models"

	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

var ev models.CmdEvent

func init() {
	f := recordCmd.Flags()
	f.StringVar(&ev.Cmd, "cmd", "", "full command line")
	f.StringVar(&ev.Shell, "shell", "zsh", "shell name (zsh/bash/fish)")
	f.StringVar(&ev.Dir, "dir", "", "working directory")
	f.StringVar(&ev.Repo, "repo", "", "git repo (optional)")
	f.StringVar(&ev.Branch, "branch", "", "git branch (optional)")
	f.Int64Var(&ev.TS, "ts", 0, "unix timestamp (seconds)")
	f.IntVar(&ev.Exit, "exit", 0, "exit code")
	f.Int64Var(&ev.Dur, "dur", 0, "duration ms (optional)")
	rootCmd.AddCommand(recordCmd)
}

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "fiks senere",
	RunE: func(cmd *cobra.Command, args []string) error {

		db, err := internal.OpenDB()

		if err != nil {
			return err
		}

		defer db.Close()

		tx, err := db.BeginTx(cmd.Context(), nil)
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(cmd.Context(),
			`INSERT INTO commands (ts,shell,dir,repo,branch,cmd,exit_code,duration_ms)
			 VALUES (?,?,?,?,?,?,?,?)`,
			ev.TS, ev.Shell, ev.Dir, ev.Repo, ev.Branch, ev.Cmd, ev.Exit, ev.Dur)

		if err != nil {
			_ = tx.Rollback()
			return err
		}

		return tx.Commit()
	},
}
