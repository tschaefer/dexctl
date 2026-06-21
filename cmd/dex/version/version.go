package version

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
)

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "Dex version",
	Run:   runVersionCmd,
}

func runVersionCmd(cmd *cobra.Command, args []string) {
	dexctl, err := cli.New(cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	version, err := dexctl.Version()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	data, err := json.Marshal(version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}
