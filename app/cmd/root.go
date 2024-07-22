package root

import (
	"fmt"
	aws "go-cli/app/cmd/aws"
	cf "go-cli/app/cmd/cf"
	gh "go-cli/app/cmd/gh"
	gl "go-cli/app/cmd/gl"
	ver "go-cli/app/cmd/ver"
	"go-cli/app/internal/config/base"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   base.NAME,
	Short: "This command-line tool allows you to create cli in go.",
	Long:  "This command-line tool allows you to create cli in go.",
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(ver.VerCmd)
	rootCmd.AddCommand(aws.AwsCmd)
	rootCmd.AddCommand(cf.CfCmd)
	rootCmd.AddCommand(gh.GhCmd)
	rootCmd.AddCommand(gl.GlCmd)
}

func Main() {
	if e := rootCmd.Execute(); e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
