package discovery

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
)

var Cmd = &cobra.Command{
	Use:   "discovery",
	Short: "Dex OIDC discovery information",
	Run:   runDiscoveryCmd,
}

func runDiscoveryCmd(cmd *cobra.Command, args []string) {
	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	discovery, err := dexctl.Discovery()
	cobra.CheckErr(err)

	data, err := json.MarshalIndent(discovery, "", "  ")
	cobra.CheckErr(err)

	fmt.Println(string(data))
}
