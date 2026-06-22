/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package user

import (
	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
	"github.com/tschaefer/dexctl/cmd/completion"
)

var deleteCmd = &cobra.Command{
	Use:               "delete",
	Short:             "Delete Dex user",
	ValidArgsFunction: completion.CompleteArgs,
	Run:               runDeleteCmd,
}

func init() {
	deleteCmd.Flags().String("user.email", "", "Email")

	_ = deleteCmd.MarkFlagRequired("user.email")
}

func runDeleteCmd(cmd *cobra.Command, args []string) {
	email, _ := cmd.Flags().GetString("user.email")

	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	err = dexctl.UserDelete(email)
	cobra.CheckErr(err)
}
