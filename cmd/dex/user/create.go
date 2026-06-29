/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package user

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
	"github.com/tschaefer/dexctl/cmd/completion"
)

var createCmd = &cobra.Command{
	Use:               "create",
	Short:             "Create Dex user",
	Run:               runCreateCmd,
	ValidArgsFunction: completion.CompleteArgs,
}

func init() {
	createCmd.Flags().StringVar(&user.Email, "user.email", "", "Email")
	createCmd.Flags().StringVar(&user.Password, "user.password", "", "Password")
	createCmd.Flags().StringVar(&user.Username, "user.username", "", "Username")
	createCmd.Flags().StringVar(&user.UserId, "user.user-id", "", "User ID")

	createCmd.Flags().String("user.config", "", "User config file (YAML or JSON)")
}

func runCreateCmd(cmd *cobra.Command, args []string) {
	err := parseUserConfig(cmd)
	cobra.CheckErr(err)

	err = verifyUserCreateConfig()
	cobra.CheckErr(err)

	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	err = dexctl.UserCreate(&user)
	cobra.CheckErr(err)
}

func verifyUserCreateConfig() error {
	if user.Email == "" || user.Password == "" || user.Username == "" || user.UserId == "" {
		return fmt.Errorf("user must have email, password, username and user-id")
	}
	return nil
}
