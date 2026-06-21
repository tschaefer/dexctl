/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package user

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/internal/dex"
)

var user dex.User

var Cmd = &cobra.Command{
	Use:   "user",
	Short: "Manage Dex users",
}

func init() {
	Cmd.AddCommand(createCmd)
	Cmd.AddCommand(deleteCmd)
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(updateCmd)
	Cmd.AddCommand(verifyPasswordCmd)
}

func parseUserConfig(cmd *cobra.Command) error {
	if !cmd.Flags().Changed("user.config") {
		return nil
	}

	configFile, _ := cmd.Flags().GetString("user.config")
	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("error reading user config file: %w", err)
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		return fmt.Errorf("error parsing user config file: %w", err)
	}

	return nil
}
