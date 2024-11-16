// Copyright (c) 2024 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cft

import (
	"context"
	"errors"
	"strings"

	"github.com/cloudflare/cloudflare-go"
)

type Client struct {
	client *cloudflare.API
	userID string
}

func NewClient(authType, token string) (*Client, error) {
	var (
		client *cloudflare.API
		err    error
	)
	switch authType {
	case "email":
		parts := strings.Split(token, ":")
		if len(parts) != 2 {
			return nil, errors.New("invalid token format: email:apikey")
		}
		client, err = cloudflare.New(parts[1], parts[0])
	default:
		client, err = cloudflare.NewWithAPIToken(token)
	}
	if err != nil {
		return nil, err
	}
	accounts, _, err := client.Accounts(context.Background(), cloudflare.AccountsListParams{})
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, errors.New("no accounts found")
	}
	return &Client{client: client, userID: accounts[0].ID}, nil
}

func (c *Client) ListTunnels() ([]cloudflare.Tunnel, error) {
	ctx := context.Background()
	tunnels, _, err := c.client.ListTunnels(
		ctx,
		cloudflare.AccountIdentifier(c.userID),
		cloudflare.TunnelListParams{
			IsDeleted: cloudflare.BoolPtr(false),
		},
	)
	if err != nil {
		return nil, err
	}
	return tunnels, nil
}

func (c *Client) DeleteTunnel(tunnelID string) error {
	ctx := context.Background()
	if err := c.client.DeleteTunnel(
		ctx,
		cloudflare.AccountIdentifier(c.userID),
		tunnelID,
	); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetTunnel(tunnelID string) (*cloudflare.Tunnel, error) {
	ctx := context.Background()
	tunnel, err := c.client.GetTunnel(
		ctx,
		cloudflare.AccountIdentifier(c.userID),
		tunnelID,
	)
	if err != nil {
		return nil, err
	}
	return &tunnel, nil
}

func (c *Client) GetTunnelConfig(tunnelID string) (*cloudflare.TunnelConfigurationResult, error) {
	ctx := context.Background()
	config, err := c.client.GetTunnelConfiguration(
		ctx,
		cloudflare.AccountIdentifier(c.userID),
		tunnelID,
	)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Client) AddTunnelIngress(tunnelID string, hostname string, service string) error {
	ctx := context.Background()
	ingress := cloudflare.UnvalidatedIngressRule{
		Hostname: hostname,
		Service:  service,
	}
	config := cloudflare.TunnelConfigurationParams{
		TunnelID: tunnelID,
		Config: cloudflare.TunnelConfiguration{
			Ingress: []cloudflare.UnvalidatedIngressRule{ingress},
		},
	}

	if _, err := c.client.UpdateTunnelConfiguration(
		ctx,
		cloudflare.AccountIdentifier(c.userID),
		config,
	); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteTunnelIngress(tunnelID string, hostname string) error {
	ctx := context.Background()
	currentConfig, err := c.GetTunnelConfig(tunnelID)
	if err != nil {
		return err
	}
	var newIngress []cloudflare.UnvalidatedIngressRule
	for _, rule := range currentConfig.Config.Ingress {
		if rule.Hostname != hostname {
			newIngress = append(newIngress, rule)
		}
	}
	config := cloudflare.TunnelConfigurationParams{
		TunnelID: tunnelID,
		Config: cloudflare.TunnelConfiguration{
			Ingress: newIngress,
		},
	}
	if _, err := c.client.UpdateTunnelConfiguration(
		ctx,
		cloudflare.AccountIdentifier(c.userID),
		config,
	); err != nil {
		return err
	}
	return nil
}
