package gl

import (
	"fmt"
	"go-cli/app/internal/config/auth"
	"go-cli/app/internal/util"
	"log"

	"github.com/spf13/cobra"
)

var GlAuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "auth",
	Long:  "auth",
	Run: func(cmd *cobra.Command, args []string) {
		gl()
	},
}

func gl() {
	auth.LoginGl()

	fmt.Printf("GITLAB_TOKEN=%s\n", auth.GitlabToken)
	fmt.Printf("GITLAB_HOST=%s\n", auth.GitlabHost)

	u, _, e := auth.Git.Users.CurrentUser()
	if e != nil {
		log.Fatal(e)
	}

	data := [][]string{
		{"Email", u.Email},
	}

	util.RenderOutput(data, []string{"Key", "Value"}, util.OutputFormat(outputFormat))
}
