// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package exec

import (
	"os"
	sysexec "os/exec"

	"github.com/ysicing/tiga/pkg/log"
)

func CommandRun(name string, args ...string) error {
	log.GetInstance().Debugf("trace exec command: %s %v", name, args)
	cmd := sysexec.Command(name, args...) // #nosec
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func Command(name string, args ...string) *sysexec.Cmd {
	log.GetInstance().Debugf("trace exec command: %s %v", name, args)
	cmd := sysexec.Command(name, args...) // #nosec
	return cmd
}
