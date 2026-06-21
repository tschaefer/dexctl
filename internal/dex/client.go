package dex

import (
	"fmt"
	"strings"

	"github.com/dexidp/dex/api/v2"
)

type Client struct {
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Secret       string   `json:"secret,omitempty"`
	RedirectUris []string `json:"redirect_uris,omitempty"`
	TrustedPeers []string `json:"trusted_peers,omitempty"`
	Public       bool     `json:"public,omitempty"`
	LogoUrl      string   `json:"logo_url,omitempty"`
}

func (d *Dex) ClientCreate(data *Client) (*Client, error) {
	request := &api.Client{
		Id:           data.Id,
		Secret:       data.Secret,
		RedirectUris: data.RedirectUris,
		TrustedPeers: data.TrustedPeers,
		Public:       data.Public,
		Name:         data.Name,
		LogoUrl:      data.LogoUrl,
	}

	response, err := d.client.CreateClient(d.ctx, &api.CreateClientReq{Client: request})
	if err != nil {
		return nil, err
	}

	if response.AlreadyExists {
		return nil, fmt.Errorf("client %s already exists", data.Id)
	}

	client := &Client{
		Id:           response.Client.Id,
		Secret:       response.Client.Secret,
		RedirectUris: response.Client.RedirectUris,
		TrustedPeers: response.Client.TrustedPeers,
		Public:       response.Client.Public,
		Name:         response.Client.Name,
		LogoUrl:      response.Client.LogoUrl,
	}

	return client, nil
}

func (d *Dex) ClientDelete(id string) error {
	response, err := d.client.DeleteClient(d.ctx, &api.DeleteClientReq{Id: id})
	if err != nil {
		return err
	}

	if response.NotFound {
		return fmt.Errorf("client %s not found", id)
	}

	return nil
}

func (d *Dex) ClientGet(id string) (*Client, error) {
	response, err := d.client.GetClient(d.ctx, &api.GetClientReq{Id: id})
	if err != nil {
		if strings.HasSuffix(err.Error(), "not found") {
			return nil, fmt.Errorf("client %s not found", id)
		}

		return nil, err
	}

	client := &Client{
		Id:           response.Client.Id,
		Secret:       response.Client.Secret,
		RedirectUris: response.Client.RedirectUris,
		TrustedPeers: response.Client.TrustedPeers,
		Public:       response.Client.Public,
		Name:         response.Client.Name,
		LogoUrl:      response.Client.LogoUrl,
	}

	return client, nil
}

func (d *Dex) ClientUpdate(data *Client) error {
	request := &api.UpdateClientReq{
		Id:           data.Id,
		RedirectUris: data.RedirectUris,
		TrustedPeers: data.TrustedPeers,
		Name:         data.Name,
		LogoUrl:      data.LogoUrl,
	}

	response, err := d.client.UpdateClient(d.ctx, request)
	if err != nil {
		return err
	}

	if response.NotFound {
		return fmt.Errorf("client %s not found", data.Id)
	}

	return nil
}

func (d *Dex) ClientList() (*[]Client, error) {
	response, err := d.client.ListClients(d.ctx, &api.ListClientReq{})
	if err != nil {
		return nil, err
	}

	var clients []Client
	for _, client := range response.Clients {
		clients = append(clients, Client{
			Id:           client.Id,
			RedirectUris: client.RedirectUris,
			TrustedPeers: client.TrustedPeers,
			Public:       client.Public,
			Name:         client.Name,
			LogoUrl:      client.LogoUrl,
		})
	}

	return &clients, nil
}
