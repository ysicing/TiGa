// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package ssh

import (
	"bytes"
	"context"
	"github.com/cockroachdb/errors"
	"golang.org/x/crypto/ssh"
	"io"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

var defaultBackoff = wait.Backoff{
	Duration: 15 * time.Second,
	Factor:   1,
	Steps:    5,
}

type SSH struct {
	SSHPort     string `json:"ssh-port,omitempty" yaml:"ssh-port,omitempty" default:"22"`
	SSHUser     string `json:"ssh-user,omitempty" yaml:"ssh-user,omitempty" default:"root"`
	SSHPassword string `json:"ssh-password,omitempty" yaml:"ssh-password,omitempty"`
	SSHKey      string `json:"ssh-key,omitempty" yaml:"ssh-key,omitempty" wrangler:"writeOnly,nullable"`
}

type Node struct {
	SSH             `json:",inline"`
	PublicIPAddress []string `json:"public-ip-address,omitempty" yaml:"public-ip-address,omitempty"`
}

// SSHDialer is a dialer that uses SSH to connect to a remote host.
type SSHDialer struct {
	sshKey     string
	sshCert    string
	sshAddress string
	username   string
	password   string
	passphrase string

	Stdin  io.ReadCloser
	Stdout io.Writer
	Stderr io.Writer
	Writer io.Writer

	Height int
	Weight int

	Term  string
	Modes ssh.TerminalModes

	ctx     context.Context
	conn    *ssh.Client
	session *ssh.Session
	cmd     *bytes.Buffer

	err error
}

func NewSSHDialer(n *Node, timeout bool) (*SSHDialer, error) {
	if len(n.PublicIPAddress) == 0 {
		return nil, errors.New("no ip address")
	}
	d := &SSHDialer{
		username:   n.SSHUser,
		password:   n.SSHPassword,
		passphrase: n.SSHPassword,
		sshKey:     n.SSHKey,
		ctx:        context.Background(),
	}
	return d, nil
}

// Dial handshake with ssh address.
func (d *SSHDialer) Dial(t bool) (*ssh.Client, error) {
	timeout := defaultBackoff.Duration
	if !t {
		timeout = 0
	}

	cfg, err := utils.GetSSHConfig(d.username, d.sshKey, d.passphrase, d.sshCert, d.password, timeout, d.useSSHAgentAuth)
	if err != nil {
		return nil, err
	}
	// establish connection with SSH server.
	return ssh.Dial("tcp", d.sshAddress, cfg)
}
