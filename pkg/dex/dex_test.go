/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package dex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDexGrpcAddr = "127.0.0.1:5557"

func Test_VersionReturnsVersion(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	assert.NoError(t, err, "create dex client")

	version, err := dex.Version()
	assert.NoError(t, err, "get version")
	assert.NotNil(t, version, "version not nil")

	assert.IsType(t, &Version{}, version, "version is type Version")
}

func Test_VersionReturnsErrorIfConnectionFails(t *testing.T) {
	dex, err := New(context.Background(), "127.0.0.1:0")
	assert.NoError(t, err, "create dex client")
	assert.NotNil(t, dex, "dex not nil")

	version, err := dex.Version()
	assert.Error(t, err, "get version")
	assert.Nil(t, version, "version nil")
}

func Test_DiscoveryReturnsDiscovery(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	assert.NoError(t, err, "create dex client")

	discovery, err := dex.Discovery()
	assert.NoError(t, err, "get discovery")
	assert.NotNil(t, discovery, "discovery not nil")

	assert.IsType(t, &Discovery{}, discovery, "discovery is type Discovery")
}

func Test_DiscoveryReturnsErrorIfConnectionFails(t *testing.T) {
	dex, err := New(context.Background(), "127.0.0.1:0")
	assert.NoError(t, err, "create dex client")
	assert.NotNil(t, dex, "dex not nil")

	discovery, err := dex.Discovery()
	assert.Error(t, err, "get discovery")
	assert.Nil(t, discovery, "discovery nil")
}
