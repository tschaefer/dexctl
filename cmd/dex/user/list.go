/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package user

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Dex users",
	Run:   runListCmd,
}

func runListCmd(cmd *cobra.Command, args []string) {
	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	response, err := dexctl.UserList()
	cobra.CheckErr(err)

	if len(*response) == 0 {
		fmt.Println("[]")
		return
	}

	data, err := json.MarshalIndent(response, "", "  ")
	cobra.CheckErr(err)

	fmt.Println(string(data))

}
