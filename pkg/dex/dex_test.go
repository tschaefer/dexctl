/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package dex

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testClientCertPath = "../../hack/dex/etc/tls/client-crt.pem"
const testClientKeyPath = "../../hack/dex/etc/tls/client-key.pem"
const testCaPath = "../../hack/dex/etc/tls/ca-crt.pem"

const testGrpcAddr = "127.0.0.1:5557"
const testNoAddr = "127.0.0.1:0"

func connectDex(t *testing.T, grpcAddr string) *Dex {
	_, tls := os.LookupEnv("DEX_TLS")

	if tls {
		dex, err := NewWithTLS(context.Background(), grpcAddr, testClientCertPath, testClientKeyPath, testCaPath)
		if err != nil {
			t.Fatal(err)
		}

		return dex
	}

	dex, err := New(context.Background(), grpcAddr)
	if err != nil {
		t.Fatal(err)
	}

	return dex
}

func Test_VersionReturnsVersion(t *testing.T) {
	dex := connectDex(t, testGrpcAddr)

	version, err := dex.Version()
	assert.NoError(t, err, "get version")
	assert.NotNil(t, version, "version not nil")

	assert.IsType(t, &Version{}, version, "version is type Version")
}

func Test_VersionReturnsErrorIfConnectionFails(t *testing.T) {
	dex := connectDex(t, testNoAddr)

	version, err := dex.Version()
	assert.Error(t, err, "get version")
	assert.Nil(t, version, "version nil")
}

func Test_DiscoveryReturnsDiscovery(t *testing.T) {
	dex := connectDex(t, testGrpcAddr)

	discovery, err := dex.Discovery()
	assert.NoError(t, err, "get discovery")
	assert.NotNil(t, discovery, "discovery not nil")

	assert.IsType(t, &Discovery{}, discovery, "discovery is type Discovery")
}

func Test_DiscoveryReturnsErrorIfConnectionFails(t *testing.T) {
	dex := connectDex(t, testNoAddr)

	discovery, err := dex.Discovery()
	assert.Error(t, err, "get discovery")
	assert.Nil(t, discovery, "discovery nil")
}
