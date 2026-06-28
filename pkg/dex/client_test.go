/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package dex

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func Test_ClientCreateReturnsClient(t *testing.T) {
	dex := connectDex(t, testGrpcAddr)

	id := gofakeit.ID()
	secret := gofakeit.UUID()
	name := gofakeit.AppName()
	redirectUris := []string{
		gofakeit.URL(),
		gofakeit.URL(),
	}

	client, err := dex.ClientCreate(&Client{
		Id:           id,
		Secret:       secret,
		Name:         name,
		RedirectUris: redirectUris,
	})
	assert.NoError(t, err, "create client")
	assert.NotNil(t, client, "client not nil")

	assert.IsType(t, &Client{}, client, "client is type Client")
	assert.Equal(t, id, client.Id, "client id")
	assert.Equal(t, secret, client.Secret, "client secret")
	assert.Equal(t, name, client.Name, "client name")
	assert.Equal(t, redirectUris, client.RedirectUris, "client redirect uris")
}

func Test_ClientCreateReturnsErrorIfConnectionFails(t *testing.T) {
	dex := connectDex(t, testNoAddr)

	client, err := dex.ClientCreate(&Client{
		Id: gofakeit.ID(),
	})
	assert.Error(t, err, "create client")
	assert.Nil(t, client, "client nil")
}

func Test_ClientCreateReturnsErrorIfClientAlreadyExists(t *testing.T) {
	dex := connectDex(t, testGrpcAddr)

	id := gofakeit.ID()
	secret := gofakeit.UUID()
	name := gofakeit.AppName()
	redirectUris := []string{
		gofakeit.URL(),
		gofakeit.URL(),
	}

	_, err := dex.ClientCreate(&Client{
		Id:           id,
		Secret:       secret,
		Name:         name,
		RedirectUris: redirectUris,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = dex.ClientCreate(&Client{
		Id:           id,
		Secret:       secret,
		Name:         name,
		RedirectUris: redirectUris,
	})
	assert.Error(t, err, "create client")
	assert.Equal(t, err.Error(), "client "+id+" already exists", "client already exists")
}

func Test_ClientDeleteReturnsClient(t *testing.T) {
	dex := connectDex(t, testGrpcAddr)

	id := gofakeit.ID()

	_, err := dex.ClientCreate(&Client{
		Id: id,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = dex.ClientDelete(id)
	assert.NoError(t, err, "delete client")
}

func Test_ClientDeleteReturnsErrorIfConnectionFails(t *testing.T) {
	dex := connectDex(t, testNoAddr)

	err := dex.ClientDelete(gofakeit.ID())
	assert.Error(t, err, "delete client")
}

func Test_ClientDeleteReturnsErrorIfClientDoesNotExist(t *testing.T) {
	dex := connectDex(t, testGrpcAddr)

	id := gofakeit.ID()

	err := dex.ClientDelete(id)
	assert.Error(t, err, "delete client")
	assert.Equal(t, err.Error(), "client "+id+" not found", "client not found")
}

func Test_ClientListReturnsClients(t *testing.T) {
	dex := connectDex(t, testGrpcAddr)

	clientCount := 10
	for range clientCount {
		_, err := dex.ClientCreate(&Client{
			Id: gofakeit.ID(),
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	clients, err := dex.ClientList()
	assert.NoError(t, err, "list clients")
	assert.NotNil(t, clients, "clients not nil")

	assert.IsType(t, &[]Client{}, clients, "clients is type []Client")
	assert.NotEmpty(t, *clients, "clients not empty")
	assert.GreaterOrEqual(t, len(*clients), clientCount, "clients count not null")
}

func Test_ClientListReturnsErrorIfConnectionFails(t *testing.T) {
	dex := connectDex(t, testNoAddr)

	clients, err := dex.ClientList()
	assert.Error(t, err, "list clients")
	assert.Nil(t, clients, "clients nil")
}

func Test_ClientGetReturnsClient(t *testing.T) {
	dex := connectDex(t, testGrpcAddr)

	id := gofakeit.ID()
	secret := gofakeit.UUID()
	name := gofakeit.AppName()
	redirectUris := []string{
		gofakeit.URL(),
		gofakeit.URL(),
	}

	_, err := dex.ClientCreate(&Client{
		Id:           id,
		Secret:       secret,
		Name:         name,
		RedirectUris: redirectUris,
	})
	if err != nil {
		t.Fatal(err)
	}

	client, err := dex.ClientGet(id)
	assert.NoError(t, err, "get client")
	assert.NotNil(t, client, "client not nil")

	assert.IsType(t, &Client{}, client, "client is type Client")
	assert.Equal(t, id, client.Id, "client id")
	assert.Equal(t, secret, client.Secret, "client secret")
	assert.Equal(t, name, client.Name, "client name")
	assert.NotEmpty(t, client.RedirectUris, "client redirect uris not empty")
	assert.Len(t, client.RedirectUris, 2, "client redirect uris count")
	assert.Equal(t, client.RedirectUris[0], redirectUris[0], "client redirect uris")
	assert.Equal(t, client.RedirectUris[1], redirectUris[1], "client redirect uris")
}

func Test_ClientGetReturnsErrorIfConnectionFails(t *testing.T) {
	dex := connectDex(t, testNoAddr)

	client, err := dex.ClientGet(gofakeit.ID())
	assert.Error(t, err, "get client")
	assert.Nil(t, client, "client nil")
}

func Test_ClientGetReturnsErrorIfClientDoesNotExist(t *testing.T) {
	dex := connectDex(t, testGrpcAddr)

	id := gofakeit.ID()

	client, err := dex.ClientGet(id)
	assert.Error(t, err, "get client")
	assert.Equal(t, err.Error(), "client "+id+" not found", "client not found")
	assert.Nil(t, client, "client nil")
}

func Test_ClientUpdateSucceeds(t *testing.T) {
	dex := connectDex(t, testGrpcAddr)

	id := gofakeit.ID()
	name := gofakeit.AppName()

	_, err := dex.ClientCreate(&Client{
		Id:   id,
		Name: name,
	})
	if err != nil {
		t.Fatal(err)
	}

	newName := gofakeit.AppName()

	err = dex.ClientUpdate(&Client{
		Id:   id,
		Name: newName,
	})
	assert.NoError(t, err)

	client, err := dex.ClientGet(id)
	assert.NoError(t, err, "get client")
	assert.NotNil(t, client, "client not nil")

	assert.NotEqual(t, name, client.Name, "client name not equal")
	assert.Equal(t, newName, client.Name, "new client name")
}

func Test_ClientUpdateReturnsErrorIfConnectionFails(t *testing.T) {
	dex := connectDex(t, testNoAddr)

	err := dex.ClientUpdate(&Client{
		Id:   gofakeit.ID(),
		Name: gofakeit.AppName(),
	})
	assert.Error(t, err, "update client")
}

func Test_ClientUpdateReturnsErrorIfClientDoesNotExist(t *testing.T) {
	dex := connectDex(t, testGrpcAddr)

	id := gofakeit.ID()

	err := dex.ClientUpdate(&Client{
		Id:   id,
		Name: gofakeit.AppName(),
	})
	assert.Error(t, err, "update client")
	assert.Equal(t, err.Error(), "client "+id+" not found", "client not found")
}
