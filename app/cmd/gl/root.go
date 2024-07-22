package gl

import (
	"github.com/spf13/cobra"
)

var outputFormat string

func init() {
	GlCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "Output format (table, raw, or json)")
}

var GlCmd = &cobra.Command{
	Use:   "gl",
	Short: "gitlab",
	Long:  "gitlab",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	GlCmd.CompletionOptions.DisableDefaultCmd = true
	GlCmd.AddCommand(GlAuthCmd)
}
