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

func userCreateSucceeds(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	err := dex.UserCreate(&User{
		Email:    gofakeit.Email(),
		Username: gofakeit.Username(),
		UserId:   gofakeit.UUID(),
		Password: gofakeit.Password(true, false, false, false, false, 32),
	})
	assert.NoError(t, err)
}

func userCreateReturnsErrorIfConnectionFails(t *testing.T) {
	dex := __connectDex(t, testNoAddr)

	err := dex.UserCreate(&User{
		Email:  gofakeit.Email(),
		UserId: gofakeit.UUID(),
	})
	assert.Error(t, err)
}

func userCreateReturnsErrorIfUserAlreadyExists(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	email := gofakeit.Email()
	username := gofakeit.Username()
	userId := gofakeit.UUID()
	password := gofakeit.Password(true, false, false, false, false, 32)

	err := dex.UserCreate(&User{
		Email:    email,
		Username: username,
		UserId:   userId,
		Password: password,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = dex.UserCreate(&User{
		Email:    email,
		Username: username,
		UserId:   userId,
		Password: password,
	})
	assert.Error(t, err, "create user")
	assert.Equal(t, err.Error(), "user "+email+" already exists")
}

func userDeleteSucceeds(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	email := gofakeit.Email()

	err := dex.UserCreate(&User{
		Email:  email,
		UserId: gofakeit.UUID(),
	})
	if err != nil {
		t.Fatal(err)
	}

	err = dex.UserDelete(email)
	assert.NoError(t, err)
}

func userDeleteReturnsErrorIfConnectionFails(t *testing.T) {
	dex := __connectDex(t, testNoAddr)

	err := dex.UserDelete(gofakeit.Email())
	assert.Error(t, err)
}

func userDeleteReturnsErrorIfUserDoesNotExist(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	email := gofakeit.Email()

	err := dex.UserDelete(email)
	assert.Error(t, err, "delete user")
	assert.Equal(t, err.Error(), "user "+email+" not found")
}

func userListReturnsUsers(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	userCount := 10
	for range userCount {
		err := dex.UserCreate(&User{
			Email:  gofakeit.Email(),
			UserId: gofakeit.UUID(),
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	users, err := dex.UserList()
	assert.NoError(t, err, "list users")
	assert.GreaterOrEqual(t, len(*users), userCount)
}

func userListReturnsErrorIfConnectionFails(t *testing.T) {
	dex := __connectDex(t, testNoAddr)

	users, err := dex.UserList()
	assert.Error(t, err)
	assert.Nil(t, users)
}

func userUpdateSucceeds(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	email := gofakeit.Email()
	username := gofakeit.Username()
	userId := gofakeit.UUID()

	err := dex.UserCreate(&User{
		Email:    email,
		UserId:   userId,
		Username: username,
	})
	if err != nil {
		t.Fatal(err)
	}

	newUsername := gofakeit.Username()

	err = dex.UserUpdate(&User{
		Email:    email,
		Username: newUsername,
	})
	assert.NoError(t, err)
}

func userUpdateReturnsErrorIfConnectionFails(t *testing.T) {
	dex := __connectDex(t, testNoAddr)

	err := dex.UserUpdate(&User{
		Email:  gofakeit.Email(),
		UserId: gofakeit.UUID(),
	})
	assert.Error(t, err)
}

func userUpdateReturnsErrorIfUserDoesNotExist(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	email := gofakeit.Email()

	err := dex.UserUpdate(&User{
		Email:  email,
		UserId: gofakeit.UUID(),
	})
	assert.Error(t, err, "update user")
	assert.Equal(t, err.Error(), "user "+email+" not found")
}

func userVerifyPasswordReturnsTrueIfPasswordIsCorrect(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	email := gofakeit.Email()
	password := gofakeit.Password(true, false, false, false, false, 32)

	err := dex.UserCreate(&User{
		Email:    email,
		Username: gofakeit.Username(),
		UserId:   gofakeit.UUID(),
		Password: password,
	})
	if err != nil {
		t.Fatal(err)
	}

	verified, err := dex.UserVerifyPassword(email, password)
	assert.True(t, verified)
	assert.NoError(t, err)
}

func userVerifyPasswordReturnsErrorIfConnectionFails(t *testing.T) {
	dex := __connectDex(t, testNoAddr)

	verified, err := dex.UserVerifyPassword(gofakeit.Email(), gofakeit.Password(true, false, false, false, false, 32))
	assert.False(t, verified)
	assert.Error(t, err)
}

func userVerifyPasswordReturnsErrorIfUserDoesNotExist(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	verified, err := dex.UserVerifyPassword(gofakeit.Email(), gofakeit.Password(true, false, false, false, false, 32))
	assert.False(t, verified)
	assert.Error(t, err)
}

func userVerifyPasswordReturnsFalseIfPasswordIsIncorrect(t *testing.T) {
	dex := __connectDex(t, testGrpcAddr)

	email := gofakeit.Email()
	password := gofakeit.Password(true, false, false, false, false, 32)

	err := dex.UserCreate(&User{
		Email:    email,
		Username: gofakeit.Username(),
		UserId:   gofakeit.UUID(),
		Password: password,
	})
	if err != nil {
		t.Fatal(err)
	}

	verified, err := dex.UserVerifyPassword(email, gofakeit.Password(true, false, false, false, false, 32))
	assert.False(t, verified)
	assert.NoError(t, err)
}

func TestDexUser(t *testing.T) {
	t.Run("dex.UserCreate successfully creates user", userCreateSucceeds)
	t.Run("dex.UserCreate returns error if Dex server is unreachable", userCreateReturnsErrorIfConnectionFails)
	t.Run("dex.UserCreate returns error if user already exists", userCreateReturnsErrorIfUserAlreadyExists)
	t.Run("dex.UserDelete successfully deletes user", userDeleteSucceeds)
	t.Run("dex.UserDelete returns error if Dex server is unreachable", userDeleteReturnsErrorIfConnectionFails)
	t.Run("dex.UserDelete returns error if user does not exist", userDeleteReturnsErrorIfUserDoesNotExist)
	t.Run("dex.UserList successfully returns user", userListReturnsUsers)
	t.Run("dex.UserList returns error if Dex server is unreachable", userListReturnsErrorIfConnectionFails)
	t.Run("dex.UserUpdate successfully updates user", userUpdateSucceeds)
	t.Run("dex.UserUpdate returns error if Dex is unreachable", userUpdateReturnsErrorIfConnectionFails)
	t.Run("dex.UserUpdate returns error if use does not exist", userUpdateReturnsErrorIfUserDoesNotExist)
	t.Run("dex.UserVerifyPassword returns true if password is correct", userVerifyPasswordReturnsTrueIfPasswordIsCorrect)
	t.Run("dex.UserVerifyPassword returns false if password is incorrect", userVerifyPasswordReturnsFalseIfPasswordIsIncorrect)
	t.Run("dex.UserVerifyPassword returns error if Dex is unreachable", userVerifyPasswordReturnsErrorIfConnectionFails)
	t.Run("dex.UserVerifyPassword returns err if user does not exist", userVerifyPasswordReturnsErrorIfUserDoesNotExist)
}
