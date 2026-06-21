package dex

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/dexidp/dex/api/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Version struct {
	Server string `json:"server,omitempty"`
	Api    int32  `json:"api,omitempty"`
}

type Discovery struct {
	Issuer                            string   `json:"issuer,omitempty"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint,omitempty"`
	TokenEndpoint                     string   `json:"token_endpoint,omitempty"`
	JwksUri                           string   `json:"jwks_uri,omitempty"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint,omitempty"`
	DeviceAuthorizationEndpoint       string   `json:"device_authorization_endpoint,omitempty"`
	IntrospectionEndpoint             string   `json:"introspection_endpoint,omitempty"`
	GrantTypesSupported               []string `json:"grant_types_supported,omitempty"`
	ResponseTypesSupported            []string `json:"response_types_supported,omitempty"`
	SubjectTypesSupported             []string `json:"subject_types_supported,omitempty"`
	IdTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported,omitempty"`
	CodeChallengeMethodsSupported     []string `json:"code_challenge_methods_supported,omitempty"`
	ScopesSupported                   []string `json:"scopes_supported,omitempty"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported,omitempty"`
	ClaimsSupported                   []string `json:"claims_supported,omitempty"`
}

type Dex struct {
	ctx    context.Context
	client api.DexClient
}

func New(ctx context.Context, grpcAddr string) (*Dex, error) {
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := api.NewDexClient(conn)

	return &Dex{
		ctx:    ctx,
		client: client,
	}, nil
}

func NewWithTLS(ctx context.Context, grpcAddr string, certPath string, keyPath string, caPath string) (*Dex, error) {
	var certPool *x509.CertPool

	if caPath != "" {
		caCert, err := os.ReadFile(caPath)
		if err != nil {
			return nil, err
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(caCert) {
			return nil, err
		}
	}

	clientCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	clientTLSConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientCert},
	}
	creds := credentials.NewTLS(clientTLSConfig)

	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to dex: %w", err)
	}
	client := api.NewDexClient(conn)

	return &Dex{
		ctx:    ctx,
		client: client,
	}, nil
}

func (d *Dex) Discovery() (*Discovery, error) {
	response, err := d.client.GetDiscovery(d.ctx, &api.DiscoveryReq{})
	if err != nil {
		return nil, err
	}

	discovery := &Discovery{
		Issuer:                            response.Issuer,
		AuthorizationEndpoint:             response.AuthorizationEndpoint,
		TokenEndpoint:                     response.TokenEndpoint,
		JwksUri:                           response.JwksUri,
		UserinfoEndpoint:                  response.UserinfoEndpoint,
		DeviceAuthorizationEndpoint:       response.DeviceAuthorizationEndpoint,
		IntrospectionEndpoint:             response.IntrospectionEndpoint,
		GrantTypesSupported:               response.GrantTypesSupported,
		ResponseTypesSupported:            response.ResponseTypesSupported,
		SubjectTypesSupported:             response.SubjectTypesSupported,
		IdTokenSigningAlgValuesSupported:  response.IdTokenSigningAlgValuesSupported,
		CodeChallengeMethodsSupported:     response.CodeChallengeMethodsSupported,
		ScopesSupported:                   response.ScopesSupported,
		TokenEndpointAuthMethodsSupported: response.TokenEndpointAuthMethodsSupported,
		ClaimsSupported:                   response.ClaimsSupported,
	}

	return discovery, nil
}

func (d *Dex) Version() (*Version, error) {
	response, err := d.client.GetVersion(d.ctx, &api.VersionReq{})
	if err != nil {
		return nil, err
	}

	version := &Version{
		Server: response.Server,
		Api:    response.Api,
	}

	return version, nil
}
