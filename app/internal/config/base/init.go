package base

import (
	"os"
)

var NAME string
var VERSION string

const defaultNAME = "go-cli"
const defaultVERSION = "dev"

func init() {

	NAME = defaultNAME

	if VERSION == "" {
		VERSION = os.Getenv("VERSION")
		if VERSION == "" {
			VERSION = defaultVERSION
		}
	}
}
