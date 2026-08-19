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

func clientCreateReturnsClient(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

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
	assert.NoError(t, err)
	assert.NotNil(t, client)

	assert.IsType(t, &Client{}, client)
	assert.Equal(t, id, client.Id)
	assert.Equal(t, secret, client.Secret)
	assert.Equal(t, name, client.Name)
	assert.Equal(t, redirectUris, client.RedirectUris)
}

func clientCreateReturnsErrorIfConnectionFails(t *testing.T) {
	dex := __connectDex(t, testNoAddr)

	client, err := dex.ClientCreate(&Client{
		Id: gofakeit.ID(),
	})
	assert.Error(t, err)
	assert.Nil(t, client)
}

func clientCreateReturnsErrorIfClientAlreadyExists(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

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
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "client "+id+" already exists")
}

func clientDeleteReturnsClient(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	id := gofakeit.ID()

	_, err := dex.ClientCreate(&Client{
		Id: id,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = dex.ClientDelete(id)
	assert.NoError(t, err)
}

func clientDeleteReturnsErrorIfConnectionFails(t *testing.T) {
	dex := __connectDex(t, testNoAddr)

	err := dex.ClientDelete(gofakeit.ID())
	assert.Error(t, err)
}

func clientDeleteReturnsErrorIfClientDoesNotExist(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	id := gofakeit.ID()

	err := dex.ClientDelete(id)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "client "+id+" not found")
}

func clientListReturnsClients(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

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
	assert.NoError(t, err)
	assert.NotNil(t, clients)

	assert.IsType(t, &[]Client{}, clients)
	assert.NotEmpty(t, *clients)
	assert.GreaterOrEqual(t, len(*clients), clientCount)
}

func clientListReturnsErrorIfConnectionFails(t *testing.T) {
	dex := __connectDex(t, testNoAddr)

	clients, err := dex.ClientList()
	assert.Error(t, err)
	assert.Nil(t, clients)
}

func clientGetReturnsClient(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

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
	assert.NoError(t, err)
	assert.NotNil(t, client)

	assert.IsType(t, &Client{}, client)
	assert.Equal(t, id, client.Id)
	assert.Equal(t, secret, client.Secret)
	assert.Equal(t, name, client.Name)
	assert.NotEmpty(t, client.RedirectUris)
	assert.Len(t, client.RedirectUris, 2)
	assert.Equal(t, client.RedirectUris[0], redirectUris[0])
	assert.Equal(t, client.RedirectUris[1], redirectUris[1])
}

func clientGetReturnsErrorIfConnectionFails(t *testing.T) {
	dex := __connectDex(t, testNoAddr)

	client, err := dex.ClientGet(gofakeit.ID())
	assert.Error(t, err)
	assert.Nil(t, client)
}

func clientGetReturnsErrorIfClientDoesNotExist(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	id := gofakeit.ID()

	client, err := dex.ClientGet(id)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "client "+id+" not found")
	assert.Nil(t, client)
}

func clientUpdateSucceeds(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

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
	assert.NoError(t, err)
	assert.NotNil(t, client)

	assert.NotEqual(t, name, client.Name)
	assert.Equal(t, newName, client.Name)
}

func clientUpdateReturnsErrorIfConnectionFails(t *testing.T) {
	dex := __connectDex(t, testNoAddr)

	err := dex.ClientUpdate(&Client{
		Id:   gofakeit.ID(),
		Name: gofakeit.AppName(),
	})
	assert.Error(t, err)
}

func clientUpdateReturnsErrorIfClientDoesNotExist(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	id := gofakeit.ID()

	err := dex.ClientUpdate(&Client{
		Id:   id,
		Name: gofakeit.AppName(),
	})
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "client "+id+" not found")
}

func TestDexClient(t *testing.T) {
	t.Run("dex.ClientCreate successfully creates client", clientCreateReturnsClient)
	t.Run("dex.ClientCreate returns error if Dex server is unreachable", clientCreateReturnsErrorIfConnectionFails)
	t.Run("dex.ClientCreate returns error if client already exist", clientCreateReturnsErrorIfClientAlreadyExists)
	t.Run("dex.ClientDelete successfully deletes client", clientDeleteReturnsClient)
	t.Run("dex.ClientDelete returns error if Dex server is unreachable", clientDeleteReturnsErrorIfConnectionFails)
	t.Run("dex.ClientDelete returns error if client does not exist", clientDeleteReturnsErrorIfClientDoesNotExist)
	t.Run("dex.ClientGet sucessfully returns client", clientGetReturnsClient)
	t.Run("dex.ClientGet returns error if Dex server is unreachable", clientGetReturnsErrorIfConnectionFails)
	t.Run("dex.ClientGet returns error if client does not exist", clientGetReturnsErrorIfClientDoesNotExist)
	t.Run("dex.ClientList successfully returns clients", clientListReturnsClients)
	t.Run("dex.ClientList returns error if Dex server is unreachable", clientListReturnsErrorIfConnectionFails)
	t.Run("dex.ClientUpdate successfully updates client", clientUpdateSucceeds)
	t.Run("dex.ClientUpdate returns error if Dex server is unreachable", clientUpdateReturnsErrorIfConnectionFails)
	t.Run("dex.ClientUpdate returns error if client does not exist", clientUpdateReturnsErrorIfClientDoesNotExist)
}
