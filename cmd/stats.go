package cmd

import (
	"fmt"
	"willbehn/what-the-terminal/internal"

	"github.com/spf13/cobra"
)

type stat struct {
	cmd      string
	countCmd int
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "A brief description of your command",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := internal.OpenDB()
		if err != nil {
			return err
		}
		defer db.Close()

		query := `SELECT cmd, count(cmd) as countCmd 
			FROM commands 
			WHERE ts >= date('now', '-7 days')
			GROUP BY cmd
			ORDER BY count(cmd) DESC
			LIMIT 10`

		rows, err := db.Query(query)
		if err != nil {
			return err
		}

		var results []stat
		for rows.Next() {
			var s stat
			if err := rows.Scan(&s.cmd, &s.countCmd); err != nil {
				return err
			}
			results = append(results, s)
		}

		fmt.Println("Top 10 Commands (Last 7 Days)")
		fmt.Println("--------------------------------")
		fmt.Printf("%-4s %-20s %s\n", "#", "Command", "Count")
		fmt.Println("--------------------------------")
		for index, row := range results {
			fmt.Printf("%-4d %-20s %d\n", index+1, row.cmd, row.countCmd)
		}
		fmt.Println("--------------------------------")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
