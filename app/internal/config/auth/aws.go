package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	AwsSession         *session.Session
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	AwsRegion          string
)

func LoginAws() {
	InitAwsSession()
}

func InitAwsSession() {
	AwsAccessKeyId, AwsSecretAccessKey, AwsRegion = os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), os.Getenv("AWS_REGION")

	var missingVars []string

	if AwsAccessKeyId == "" {
		missingVars = append(missingVars, "AWS_ACCESS_KEY_ID")
	}
	if AwsSecretAccessKey == "" {
		missingVars = append(missingVars, "AWS_SECRET_ACCESS_KEY")
	}
	if AwsRegion == "" {
		missingVars = append(missingVars, "AWS_REGION")
	}

	if len(missingVars) > 0 {
		e := fmt.Sprintf("Missing environment variables: %v", missingVars)
		log.Fatal(e)
	}

	if AwsSession == nil {
		var e error
		AwsSession, e = session.NewSession(&aws.Config{
			Region:      aws.String(AwsRegion),
			Credentials: credentials.NewStaticCredentials(AwsAccessKeyId, AwsSecretAccessKey, ""),
		})
		if e != nil {
			log.Fatal(e)
		}
	}
}
