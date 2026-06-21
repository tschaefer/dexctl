/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package discovery

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
	"github.com/tschaefer/dexctl/cmd/completion"
)

var Cmd = &cobra.Command{
	Use:               "discovery",
	Short:             "Dex OIDC discovery information",
	Run:               runDiscoveryCmd,
	ValidArgsFunction: completion.CompleteArgs,
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
