package dex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDexGrpcAddr = "localhost:5557"

func Test_VersionReturnsVersion(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	assert.NoError(t, err, "create dex client")

	version, err := dex.Version()
	assert.NoError(t, err, "get version")
	assert.NotNil(t, version, "version not nil")

	assert.IsType(t, &Version{}, version, "version is type Version")
}

func Test_VersionReturnsErrorIfConnectionFails(t *testing.T) {
	dex, err := New(context.Background(), "localhost:0")
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
	dex, err := New(context.Background(), "localhost:0")
	assert.NoError(t, err, "create dex client")
	assert.NotNil(t, dex, "dex not nil")

	discovery, err := dex.Discovery()
	assert.Error(t, err, "get discovery")
	assert.Nil(t, discovery, "discovery nil")
}
