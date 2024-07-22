package gh

import (
	"context"
	"fmt"
	"go-cli/app/internal/config/auth"
	"go-cli/app/internal/util"
	"log"

	"github.com/spf13/cobra"
)

var GhAuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "auth",
	Long:  "auth",
	Run: func(cmd *cobra.Command, args []string) {
		gh()
	},
}

func gh() {
	auth.LoginGh()

	fmt.Printf("GITHUB_TOKEN=%s\n", auth.GithubToken)

	u, _, e := auth.Client.Users.Get(context.Background(), "")
	if e != nil {
		log.Fatal(e)
	}

	data := [][]string{
		{"Email", *u.Email},
	}

	util.RenderOutput(data, []string{"Key", "Value"}, util.OutputFormat(outputFormat))
}
