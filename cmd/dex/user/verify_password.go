package user

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
)

var verifyPasswordCmd = &cobra.Command{
	Use:   "verify-password",
	Short: "Verify Dex user password",
	Args:  cobra.ExactArgs(2),
	Run:   runVerifyPasswordCmd,

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
