package aws

import (
	"fmt"
	"go-cli/app/internal/config/auth"
	"go-cli/app/internal/util"
	"log"

	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/spf13/cobra"
)

var AwsAuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "auth",
	Long:  "auth",
	Run: func(cmd *cobra.Command, args []string) {
		aws()
	},
}

func aws() {
	auth.LoginAws()
	svc := sts.New(auth.AwsSession)

	fmt.Printf("AWS_ACCESS_KEY_ID=%s\n", auth.AwsAccessKeyId)
	fmt.Printf("AWS_SECRET_ACCESS_KEY=%s\n", auth.AwsSecretAccessKey)
	fmt.Printf("AWS_DEFAULT_REGION=%s\n", auth.AwsRegion)

	u, e := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if e != nil {
		log.Panic(e)
	}

	data := [][]string{
		{"Account", *u.Account},
	}

	util.RenderOutput(data, []string{"Key", "Value"}, util.OutputFormat(outputFormat))
}
