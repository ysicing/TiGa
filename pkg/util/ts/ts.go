// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package ts

import (
	"context"
	"time"

	"github.com/ysicing/tiga/pkg/log"
	"tailscale.com/client/tailscale"
	"tailscale.com/ipn/ipnstate"
	"tailscale.com/types/key"
)

func findActiveExitNodeFromPeersMap(peers map[key.NodePublic]*ipnstate.PeerStatus) *ipnstate.PeerStatus {
	for _, p := range peers {
		if p.ExitNode {
			return p
		}
	}
	return nil
}

// GetTailscaleStatus returns a string describing the tailscale status
func GetTailscaleStatus() {
	logpkg := log.GetInstance()
	// check tailscale status
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ts, err := tailscale.Status(ctx) // https://pkg.go.dev/tailscale.com@v1.40.0/ipn/ipnstate#Status
	defer cancel()

	if err != nil {
		logpkg.Warnf("determine tailscaled status failed: %s", err)
		return
	}

	// https://github.com/tailscale/tailscale/blob/9bdaece3d7c3c83aae01e0736ba54e833f4aea51/cmd/tailscale/cli/status.go#L162-L196

	if !ts.Self.Online {
		logpkg.Warnf("tailscale failed connected to tsnet: BackendState = %s", ts.BackendState)
		return
	}

	exitNodeStatus := findActiveExitNodeFromPeersMap(ts.Peer)

	if exitNodeStatus == nil {
		logpkg.Donef("tailscale online on tsnet and not using any exit node")
		return
	}
	if exitNodeStatus.Active {
		if exitNodeStatus.Relay != "" && exitNodeStatus.CurAddr == "" {
			logpkg.Donef("tailscale online on tsnet, exit node %s via relay %s", exitNodeStatus.HostName, exitNodeStatus.Relay)
			return
		}

		if exitNodeStatus.CurAddr != "" {
			logpkg.Donef("tailscale online on tsnet, exit node %s via %s", exitNodeStatus.HostName, exitNodeStatus.CurAddr)
			return
		}
		logpkg.Warnf("tailscale online on tsnet, exit node %s (unknown connection)", exitNodeStatus.HostName)
		return
	}
	logpkg.Warnf("tailscale online on tsnet, exit node %s is inactive", exitNodeStatus.HostName)
}
