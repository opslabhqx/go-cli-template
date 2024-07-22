package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

var (
	Api                *cloudflare.API
	CloudflareApiKey   string
	CloudflareApiEmail string
)

func LoginCf() {
	InitCfSession()
}

func InitCfSession() {
	CloudflareApiKey, CloudflareApiEmail = os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL")

	var missingVars []string

	if CloudflareApiKey == "" {
		missingVars = append(missingVars, "CLOUDFLARE_API_KEY")
	}
	if CloudflareApiEmail == "" {
		missingVars = append(missingVars, "CLOUDFLARE_API_EMAIL")
	}

	if len(missingVars) > 0 {
		e := fmt.Sprintf("Missing environment variables: %v", missingVars)
		log.Fatal(e)
	}

	if Api == nil {
		var e error
		Api, e = cloudflare.New(CloudflareApiKey, CloudflareApiEmail)
		if e != nil {
			log.Fatal(e)
		}
	}
}
