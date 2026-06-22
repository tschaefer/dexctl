/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package dex

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func Test_UserCreateSucceeds(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	if err != nil {
		t.Fatal(err)
	}

	err = dex.UserCreate(&User{
		Email:    gofakeit.Email(),
		Username: gofakeit.Username(),
		UserId:   gofakeit.UUID(),
		Password: gofakeit.Password(true, false, false, false, false, 32),
	})
	assert.NoError(t, err, "create user")
}

func Test_UserCreateReturnsErrorIfConnectionFails(t *testing.T) {
	dex, err := New(context.Background(), "localhost:0")
	if err != nil {
		t.Fatal(err)
	}

	err = dex.UserCreate(&User{
		Email:  gofakeit.Email(),
		UserId: gofakeit.UUID(),
	})
	assert.Error(t, err, "create user")
}

func Test_UserCreateReturnsErrorIfUserAlreadyExists(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	if err != nil {
		t.Fatal(err)
	}

	email := gofakeit.Email()
	username := gofakeit.Username()
	userId := gofakeit.UUID()
	password := gofakeit.Password(true, false, false, false, false, 32)

	err = dex.UserCreate(&User{
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
	assert.Equal(t, err.Error(), "user "+email+" already exists", "user already exists")
}

func Test_UserDeleteSucceeds(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	if err != nil {
		t.Fatal(err)
	}

	email := gofakeit.Email()

	err = dex.UserCreate(&User{
		Email:  email,
		UserId: gofakeit.UUID(),
	})
	if err != nil {
		t.Fatal(err)
	}

	err = dex.UserDelete(email)
	assert.NoError(t, err, "delete user")
}

func Test_UserDeleteReturnsErrorIfConnectionFails(t *testing.T) {
	dex, err := New(context.Background(), "localhost:0")
	if err != nil {
		t.Fatal(err)
	}

	err = dex.UserDelete(gofakeit.Email())
	assert.Error(t, err, "delete user")
}

func Test_UserDeleteReturnsErrorIfUserDoesNotExist(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	if err != nil {
		t.Fatal(err)
	}

	email := gofakeit.Email()

	err = dex.UserDelete(email)
	assert.Error(t, err, "delete user")
	assert.Equal(t, err.Error(), "user "+email+" not found", "user not found")
}

func Test_UserListReturnsUsers(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	if err != nil {
		t.Fatal(err)
	}

	userCount := 10
	for range userCount {
		err = dex.UserCreate(&User{
			Email:  gofakeit.Email(),
			UserId: gofakeit.UUID(),
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	users, err := dex.UserList()
	assert.NoError(t, err, "list users")
	assert.GreaterOrEqual(t, len(*users), userCount, "users count not null")
}

func Test_UserListReturnsErrorIfConnectionFails(t *testing.T) {
	dex, err := New(context.Background(), "localhost:0")
	if err != nil {
		t.Fatal(err)
	}

	users, err := dex.UserList()
	assert.Error(t, err, "list users")
	assert.Nil(t, users, "users nil")
}

func Test_UserUpdateSucceeds(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	if err != nil {
		t.Fatal(err)
	}

	email := gofakeit.Email()
	username := gofakeit.Username()
	userId := gofakeit.UUID()

	err = dex.UserCreate(&User{
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
	assert.NoError(t, err, "update user")
}

func Test_UserUpdateReturnsErrorIfConnectionFails(t *testing.T) {
	dex, err := New(context.Background(), "localhost:0")
	if err != nil {
		t.Fatal(err)
	}

	err = dex.UserUpdate(&User{
		Email:  gofakeit.Email(),
		UserId: gofakeit.UUID(),
	})
	assert.Error(t, err, "update user")
}

func Test_UserUpdateReturnsErrorIfUserDoesNotExist(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	if err != nil {
		t.Fatal(err)
	}

	email := gofakeit.Email()

	err = dex.UserUpdate(&User{
		Email:  email,
		UserId: gofakeit.UUID(),
	})
	assert.Error(t, err, "update user")
	assert.Equal(t, err.Error(), "user "+email+" not found", "user not found")
}

func Test_UserVerifyPasswordReturnsFalseIfPasswordIsCorrect(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	if err != nil {
		t.Fatal(err)
	}

	email := gofakeit.Email()
	password := gofakeit.Password(true, false, false, false, false, 32)

	err = dex.UserCreate(&User{
		Email:    email,
		Username: gofakeit.Username(),
		UserId:   gofakeit.UUID(),
		Password: password,
	})
	if err != nil {
		t.Fatal(err)
	}

	verified, err := dex.UserVerifyPassword(email, password)
	assert.True(t, verified, "verify password")
	assert.NoError(t, err, "verify password")
}

func Test_UserVerifyPasswordReturnsErrorIfConnectionFails(t *testing.T) {
	dex, err := New(context.Background(), "localhost:0")
	if err != nil {
		t.Fatal(err)
	}

	verified, err := dex.UserVerifyPassword(gofakeit.Email(), gofakeit.Password(true, false, false, false, false, 32))
	assert.False(t, verified, "verify password")
	assert.Error(t, err, "verify password")
}

func Test_UserVerifyPasswordReturnsErrorIfUserDoesNotExist(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	if err != nil {
		t.Fatal(err)
	}

	verified, err := dex.UserVerifyPassword(gofakeit.Email(), gofakeit.Password(true, false, false, false, false, 32))
	assert.False(t, verified, "verify password")
	assert.Error(t, err, "verify password")
}

func Test_UserVerifyPasswordReturnsFalseIfPasswordIsIncorrect(t *testing.T) {
	dex, err := New(context.Background(), testDexGrpcAddr)
	if err != nil {
		t.Fatal(err)
	}

	email := gofakeit.Email()
	password := gofakeit.Password(true, false, false, false, false, 32)

	err = dex.UserCreate(&User{
		Email:    email,
		Username: gofakeit.Username(),
		UserId:   gofakeit.UUID(),
		Password: password,
	})
	if err != nil {
		t.Fatal(err)
	}

	verified, err := dex.UserVerifyPassword(email, gofakeit.Password(true, false, false, false, false, 32))
	assert.False(t, verified, "verify password")
	assert.NoError(t, err, "verify password")
}
