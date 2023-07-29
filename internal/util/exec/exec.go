// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package exec

import (
	"os"
	sysexec "os/exec"
	"strings"

	"github.com/ergoapi/util/file"
	"github.com/ysicing/tiga/common"
)

func trace(cmd *sysexec.Cmd) {
	key := strings.Join(cmd.Args, " ")
	file.WriteFile(common.GetDefaultLogFile("exec.trace.log"), key, false)
}

func Command(name string, arg ...string) *sysexec.Cmd {
	cmd := sysexec.Command(name, arg...) // #nosec
	// cmd.Dir = common.GetDefaultCacheDir()
	trace(cmd)
	return cmd
}

func CommandRun(name string, arg ...string) error {
	cmd := sysexec.Command(name, arg...) // #nosec
	trace(cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	// cmd.Dir = common.GetDefaultCacheDir()
	return cmd.Run()
}

func CommandBashRunWithResp(cmdStr string) (string, error) {
	cmd := sysexec.Command("/bin/bash", "-c", cmdStr) // #nosec
	// cmd.Dir = common.GetDefaultCacheDir()
	trace(cmd)
	result, err := cmd.CombinedOutput()
	return string(result), err
}
