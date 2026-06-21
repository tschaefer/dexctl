package client

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Dex password",
	Run:   runCreateCmd,
}

func init() {
	createCmd.Flags().StringVar(&client.Id, "client.id", "", "Unique Identifier")
	createCmd.Flags().StringVar(&client.Secret, "client.secret", "", "Secret")
	createCmd.Flags().StringSliceVar(&client.RedirectUris, "client.redirect-uris", []string{}, "Redirect URIs")
	createCmd.Flags().StringSliceVar(&client.TrustedPeers, "client.trusted-peers", []string{}, "Trusted Peers")
	createCmd.Flags().BoolVar(&client.Public, "client.public", false, "Public")
	createCmd.Flags().StringVar(&client.Name, "client.name", "", "Name")
	createCmd.Flags().StringVar(&client.LogoUrl, "client.logo-url", "", "Logo URL")

	createCmd.Flags().String("client.config", "", "Client config file")
}

func runCreateCmd(cmd *cobra.Command, args []string) {
	err := parseClientConfig(cmd)
	cobra.CheckErr(err)

	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	response, err := dexctl.ClientCreate(&client)
	cobra.CheckErr(err)

	data, err := json.MarshalIndent(response, "", "  ")
	cobra.CheckErr(err)

	fmt.Println(string(data))
}
