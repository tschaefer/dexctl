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

var updateCmd = &cobra.Command{
	Use:               "update",
	Short:             "Update Dex clients",
	Run:               runUpdateCmd,
	ValidArgsFunction: completion.CompleteArgs,
}

func init() {
	updateCmd.Flags().StringVar(&client.Id, "client.id", "", "Unique Identifier")
	updateCmd.Flags().StringSliceVar(&client.RedirectUris, "client.redirect-uris", []string{}, "Redirect URIs")
	updateCmd.Flags().StringSliceVar(&client.TrustedPeers, "client.trusted-peers", []string{}, "Trusted Peers")
	updateCmd.Flags().StringVar(&client.Name, "client.name", "", "Name")
	updateCmd.Flags().StringVar(&client.LogoUrl, "client.logo-url", "", "Logo URL")

	updateCmd.Flags().String("client.config", "", "Client config file (YAML or JSON)")
}

func runUpdateCmd(cmd *cobra.Command, args []string) {
	err := parseClientConfig(cmd)
	cobra.CheckErr(err)

	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	err = dexctl.ClientUpdate(&client)
	cobra.CheckErr(err)
}
