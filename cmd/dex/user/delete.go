/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package user

import (
	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Dex user",
	Args:  cobra.ExactArgs(1),
	Run:   runDeleteCmd,
}

func runDeleteCmd(cmd *cobra.Command, args []string) {
	email := args[0]

	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	err = dexctl.UserDelete(email)
	cobra.CheckErr(err)
}
