/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package version

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
	"github.com/tschaefer/dexctl/cmd/completion"
)

var Cmd = &cobra.Command{
	Use:               "version",
	Short:             "Dex version",
	Run:               runVersionCmd,
	ValidArgsFunction: completion.CompleteArgs,
}

func runVersionCmd(cmd *cobra.Command, args []string) {
	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	version, err := dexctl.Version()
	cobra.CheckErr(err)

	data, err := json.Marshal(version)
	cobra.CheckErr(err)

	fmt.Println(string(data))
}
