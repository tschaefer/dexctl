/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package client

import (
	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
	"github.com/tschaefer/dexctl/cmd/completion"
)

var deleteCmd = &cobra.Command{
	Use:               "delete",
	Short:             "Delete Dex clients",
	ValidArgsFunction: completion.CompleteArgs,
	Run:               runDeleteCmd,
}

func init() {
	getCmd.Flags().String("client.id", "", "Client ID")

	_ = getCmd.MarkFlagRequired("client.id")
}

func runDeleteCmd(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetString("client.id")

	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	err = dexctl.ClientDelete(id)
	cobra.CheckErr(err)
}
