package gh

import (
	"github.com/spf13/cobra"
)

var outputFormat string

func init() {
	GhCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "Output format (table, raw, or json)")
}

var GhCmd = &cobra.Command{
	Use:   "gh",
	Short: "github",
	Long:  "github",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	GhCmd.CompletionOptions.DisableDefaultCmd = true
	GhCmd.AddCommand(GhAuthCmd)
}
