/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package dex

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/dexidp/dex/api/v2"
)

type User struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
	UserId   string `json:"user_id,omitempty"`
}

func (d *Dex) UserCreate(data *User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	request := &api.Password{
		Email:    data.Email,
		Hash:     hash,
		Username: data.Username,
		UserId:   data.UserId,
	}

	response, err := d.client.CreatePassword(d.ctx, &api.CreatePasswordReq{Password: request})
	if err != nil {
		return err
	}

	if response.AlreadyExists {
		return fmt.Errorf("user %s already exists", data.Email)
	}

	return nil
}

func (d *Dex) UserDelete(email string) error {
	response, err := d.client.DeletePassword(d.ctx, &api.DeletePasswordReq{Email: email})
	if err != nil {
		return err
	}

	if response.NotFound {
		return fmt.Errorf("user %s not found", email)
	}

	return nil
}

func (d *Dex) UserList() (*[]User, error) {
	response, err := d.client.ListPasswords(d.ctx, &api.ListPasswordReq{})
	if err != nil {
		return nil, err
	}

	var passwords []User
	for _, password := range response.Passwords {
		passwords = append(passwords, User{
			Email:    password.Email,
			Username: password.Username,
			UserId:   password.UserId,
			Password: string(password.Hash),
		})
	}

	return &passwords, nil
}

func (d *Dex) UserUpdate(data *User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	client := &api.UpdatePasswordReq{
		Email:       data.Email,
		NewHash:     hash,
		NewUsername: data.Username,
	}

	response, err := d.client.UpdatePassword(d.ctx, client)
	if err != nil {
		return err
	}

	if response.NotFound {
		return fmt.Errorf("user %s not found", data.Email)
	}

	return nil
}

func (d *Dex) UserVerifyPassword(email string, password string) (bool, error) {
	response, err := d.client.VerifyPassword(d.ctx, &api.VerifyPasswordReq{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return false, err
	}

	if response.NotFound {
		return false, fmt.Errorf("user %s not found", email)
	}

	if response.Verified {
		return true, nil
	}

	return false, nil
}
