package user

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/cli"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Dex user",
	Run:   runUpdateCmd,
}

func init() {
	updateCmd.Flags().StringVar(&user.Email, "user.email", "", "Email")
	updateCmd.Flags().StringVar(&user.Password, "user.password", "", "Password")
	updateCmd.Flags().StringVar(&user.Username, "user.username", "", "Username")

	updateCmd.Flags().String("user.config", "", "User config file")
}

func runUpdateCmd(cmd *cobra.Command, args []string) {
	err := parseUserConfig(cmd)
	cobra.CheckErr(err)

	err = verifyUserUpdateConfig()
	cobra.CheckErr(err)

	dexctl, err := cli.New(cmd)
	cobra.CheckErr(err)

	err = dexctl.UserUpdate(&user)
	cobra.CheckErr(err)
}

func verifyUserUpdateConfig() error {
	if user.Email == "" {
		return fmt.Errorf("user must have email")
	}
	return nil
}
