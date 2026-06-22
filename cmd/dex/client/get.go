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
	"github.com/tschaefer/dexctl/cmd/completion"
)

var getCmd = &cobra.Command{
	Use:               "get",
	Short:             "Get Dex clients",
	ValidArgsFunction: completion.CompleteArgs,
	Run:               runGetCmd,
}

func init() {
	getCmd.Flags().String("client.id", "", "Client ID")

	_ = getCmd.MarkFlagRequired("client.id")
}

func runGetCmd(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetString("client.id")

	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	response, err := dexctl.ClientGet(id)
	cobra.CheckErr(err)

	data, err := json.MarshalIndent(response, "", "  ")
	cobra.CheckErr(err)

	fmt.Println(string(data))

}
