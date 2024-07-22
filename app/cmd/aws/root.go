package aws

import (
	"github.com/spf13/cobra"
)

var outputFormat string

func init() {
	AwsAuthCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "Output format (table, raw, or json)")
}

var AwsCmd = &cobra.Command{
	Use:   "aws",
	Short: "amazon web services",
	Long:  "amazon web services",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	AwsCmd.CompletionOptions.DisableDefaultCmd = true
	AwsCmd.AddCommand(AwsAuthCmd)
}
