package cli

import (
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/tschaefer/dexctl/internal/dex"
)

type config struct {
	GprcAddr string
	WithTLS  bool
	CertPath string
	KeyPath  string
	CaPath   string
}

func New(cmd *cobra.Command) (*dex.Dex, error) {
	config, err := newConfig()
	if err != nil {
		return nil, err
	}

	if config.WithTLS {
		dexctl, err := dex.NewWithTLS(cmd.Context(), config.GprcAddr, config.CertPath, config.KeyPath, config.CaPath)
		if err != nil {
			return nil, err
		}
		return dexctl, nil
	}

	dexctl, err := dex.New(cmd.Context(), config.GprcAddr)
	if err != nil {
		return nil, err
	}
	return dexctl, nil
}

func newConfig() (*config, error) {
	grpcAddr, hasGrpcAddr := os.LookupEnv("DEXCTL_GRPC_ADDRESS")
	if !hasGrpcAddr {
		return nil, errors.New("missing dex grpc address")
	}

	certPath, hasCertPath := os.LookupEnv("DEXCTL_CERT_PATH")
	keyPath, hasKeyPath := os.LookupEnv("DEXCTL_KEY_PATH")

	if hasCertPath && !hasKeyPath {
		return nil, errors.New("missing key path")
	}

	if hasKeyPath && !hasCertPath {
		return nil, errors.New("missing cert path")
	}

	withTLS := hasCertPath && hasKeyPath

	caPath := os.Getenv("DEXCTL_CA_PATH")

	return &config{
		GprcAddr: grpcAddr,
		WithTLS:  withTLS,
		CertPath: certPath,
		KeyPath:  keyPath,
		CaPath:   caPath,
	}, nil
}
