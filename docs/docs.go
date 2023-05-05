// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package main

import (
	"strings"

	"github.com/ergoapi/util/file"
	"github.com/ergoapi/util/github"
	"github.com/ergoapi/util/version"
	"github.com/spf13/cobra/doc"
	"github.com/ysicing/tiga/cmd"
	"github.com/ysicing/tiga/pkg/factory"
)

func main() {
	f := factory.DefaultFactory()
	tcli := cmd.BuildRoot(f)
	err := doc.GenMarkdownTree(tcli, "./docs")
	if err != nil {
		panic(err)
	}
	pkg := github.Pkg{
		Owner: "ysicing",
		Repo:  "tiga",
	}
	tag, err := pkg.LastTag()
	if err != nil {
		return
	}
	file.WriteFile("VERSION", strings.TrimPrefix(version.Next(tag.Name, false, false, true), "v"), true)
}
