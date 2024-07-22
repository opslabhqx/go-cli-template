package cf

import (
	"github.com/spf13/cobra"
)

var outputFormat string

func init() {
	CfCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "Output format (table, raw, or json)")
}

var CfCmd = &cobra.Command{
	Use:   "cf",
	Short: "cloudflare",
	Long:  "cloudflare",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	CfCmd.CompletionOptions.DisableDefaultCmd = true
	CfCmd.AddCommand(CfAuthCmd)
}
