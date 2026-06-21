/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package user

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
	"github.com/tschaefer/dexctl/cmd/completion"
)

var verifyPasswordCmd = &cobra.Command{
	Use:               "verify-password",
	Short:             "Verify Dex user password",
	ValidArgsFunction: completion.CompleteArgs,
	Run:               runVerifyPasswordCmd,
}

func init() {
	verifyPasswordCmd.Flags().String("user.email", "", "Email")
	verifyPasswordCmd.Flags().String("user.password", "", "Password")

	_ = deleteCmd.MarkFlagRequired("user.email")
	_ = deleteCmd.MarkFlagRequired("user.password")
}

func runVerifyPasswordCmd(cmd *cobra.Command, args []string) {
	email, _ := cmd.Flags().GetString("user.email")
	password, _ := cmd.Flags().GetString("user.password")

	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	verified, err := dexctl.UserVerifyPassword(email, password)
	cobra.CheckErr(err)

	if !verified {
		os.Exit(1)
	}
}
