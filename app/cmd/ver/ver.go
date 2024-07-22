package ver

import (
	"fmt"
	"go-cli/app/internal/config/base"
	"runtime"

	"github.com/spf13/cobra"
)

var VerCmd = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Print the version number of %s", base.NAME),
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("%s %s %s\n", base.NAME, base.VERSION, runtime.GOARCH)
	},
}
