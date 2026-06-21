/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package client

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/pkg/dex"
)

var client dex.Client

var Cmd = &cobra.Command{
	Use:   "client",
	Short: "Manage Dex clients",
}

func init() {
	Cmd.AddCommand(createCmd)
	Cmd.AddCommand(deleteCmd)
	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(updateCmd)
}

func parseClientConfig(cmd *cobra.Command) error {
	if !cmd.Flags().Changed("client.config") {
		return nil
	}

	configFile, _ := cmd.Flags().GetString("client.config")
	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("error reading client config file: %w", err)
	}

	err = json.Unmarshal(data, &client)
	if err != nil {
		return fmt.Errorf("error parsing client config file: %w", err)
	}

	return nil
}
