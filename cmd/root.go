/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/cmd/dex/client"
	"github.com/tschaefer/dexctl/cmd/dex/discovery"
	"github.com/tschaefer/dexctl/cmd/dex/user"
	"github.com/tschaefer/dexctl/cmd/dex/version"
)

var rootCmd = &cobra.Command{
	Use:              "dexctl",
	Short:            "Command line tool for managing a dex server",
	Long:             `Command line tool for managing a dex server`,
	Run:              runRootCmd,
	PersistentPreRun: handlePersistentFlags,
}

func Execute() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %v\n", r)
			os.Exit(1)
		}
	}()

	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "version")

	rootCmd.PersistentFlags().String("grpc-address", "", "dex grpc address")
	rootCmd.PersistentFlags().String("certificate-path", "", "certificate to use for grpc requests")
	rootCmd.PersistentFlags().String("key-path", "", "key to use for grpc requests")
	rootCmd.PersistentFlags().String("ca-path", "", "ca to use for grpc requests")

	rootCmd.AddCommand(client.Cmd)
	rootCmd.AddCommand(discovery.Cmd)
	rootCmd.AddCommand(user.Cmd)
	rootCmd.AddCommand(version.Cmd)
}

func handlePersistentFlags(cmd *cobra.Command, args []string) {
	dexGrpcAddress, _ := cmd.Flags().GetString("grpc-address")
	dexCertificatePath, _ := cmd.Flags().GetString("certificate-path")
	dexKeyPath, _ := cmd.Flags().GetString("key-path")
	dexCaPath, _ := cmd.Flags().GetString("ca-path")

	envs := map[string]string{
		"DEXCTL_GRPC_ADDRESS":     dexGrpcAddress,
		"DEXCTL_CERTIFICATE_PATH": dexCertificatePath,
		"DEXCTL_KEY_PATH":         dexKeyPath,
		"DEXCTL_CA_PATH":          dexCaPath,
	}

	for k, v := range envs {
		if v != "" {
			if err := os.Setenv(k, v); err != nil {
				panic(err)
			}
		}
	}
}

func runRootCmd(cmd *cobra.Command, args []string) {
	version, _ := cmd.Flags().GetBool("version")
	if version {
		fmt.Println("dexctl version 0.0.1")
		return
	}
	_ = cmd.Help()
}
