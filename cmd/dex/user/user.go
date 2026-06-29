/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/pkg/dex"
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

	if jsonerr := json.Unmarshal(data, &user); jsonerr != nil {
		if yamlerr := yaml.Unmarshal(data, &user); yamlerr != nil {
			return fmt.Errorf("error parsing user config file: %w", errors.Join(jsonerr, yamlerr))
		}
	}

	return nil
}
