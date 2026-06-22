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
	Args:              cobra.ExactArgs(2),
	ValidArgsFunction: completion.CompleteArgs,
	Run:               runVerifyPasswordCmd,
}

func runVerifyPasswordCmd(cmd *cobra.Command, args []string) {
	email := args[0]
	password := args[1]

	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	verified, err := dexctl.UserVerifyPassword(email, password)
	cobra.CheckErr(err)

	if !verified {
		os.Exit(1)
	}
}
