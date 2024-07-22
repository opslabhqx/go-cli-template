package cf

import (
	"context"
	"fmt"
	"go-cli/app/internal/config/auth"
	"go-cli/app/internal/util"
	"log"

	"github.com/spf13/cobra"
)

var CfAuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "auth",
	Long:  "auth",
	Run: func(cmd *cobra.Command, args []string) {
		cf()
	},
}

func cf() {
	auth.LoginCf()
	ctx := context.Background()

	fmt.Printf("CLOUDFLARE_API_KEY=%s\n", auth.CloudflareApiKey)
	fmt.Printf("CLOUDFLARE_API_EMAIL=%s\n", auth.CloudflareApiEmail)

	u, e := auth.Api.UserDetails(ctx)
	if e != nil {
		log.Fatal(e)
	}

	data := [][]string{
		{"Email", u.Email},
	}

	util.RenderOutput(data, []string{"Key", "Value"}, util.OutputFormat(outputFormat))
}
