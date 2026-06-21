/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package client

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Dex clients",
	Args:  cobra.ExactArgs(1),
	Run:   runGetCmd,
}

func runGetCmd(cmd *cobra.Command, args []string) {
	id := args[0]

	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	response, err := dexctl.ClientGet(id)
	cobra.CheckErr(err)

	data, err := json.MarshalIndent(response, "", "  ")
	cobra.CheckErr(err)

	fmt.Println(string(data))

}
