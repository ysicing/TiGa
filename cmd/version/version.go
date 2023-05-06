// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package version

import (
	"fmt"
	"html/template"
	"os"
	"runtime"
	"strings"
	"text/tabwriter"

	gv "github.com/Masterminds/semver/v3"
	"github.com/cockroachdb/errors"
	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/github"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/pkg/factory"
	logpkg "github.com/ysicing/tiga/pkg/log"
)

var versionTpl = `{{with .Client -}}
Client:
 Version:           {{ .Version }}
 Go version:        {{ .GoVersion }}
 Git commit:        {{ .GitCommit }}
 Built:             {{ .BuildTime }}
 OS/Arch:           {{.Os}}/{{.Arch}}
 Experimental:      {{.Experimental}}
{{- if .CanUpgrade }}
 Note:              {{ .UpgradeMessage }}
 URL:               https://github.com/easysoft/quickon_cli/releases/tag/v{{ .LastVersion }}
{{- end }}
{{- end}}
`

const (
	defaultVersion       = "0.0.0"
	defaultGitCommitHash = "a1b2c3d4"
	defaultBuildDate     = "Mon Aug  3 15:06:50 2020"
)

type versionInfo struct {
	Client clientVersion
}

type clientVersion struct {
	Version        string
	LastVersion    string
	GoVersion      string
	GitCommit      string
	Os             string
	Arch           string
	BuildTime      string `json:",omitempty"`
	Experimental   bool
	CanUpgrade     bool
	UpgradeMessage string
}

// PreCheckLatestVersion 检查最新版本
func PreCheckLatestVersion(log logpkg.Logger) (version, t string, err error) {
	version, _ = checkLastVersionFromGithub()
	if version != "" {
		log.Debugf("fetch version from github: %s", version)
		return version, "github", nil
	}
	version, err = checkLatestVersionFromAPI()
	if err != nil {
		return version, "api", err
	}
	log.Debugf("fetch version from api: %s", version)
	return version, "api", nil
}

func checkLastVersionFromGithub() (string, error) {
	pkg := github.Pkg{
		Owner: "ysicing",
		Repo:  "tiga",
	}
	tag, err := pkg.LastTag()
	if err != nil {
		return "", err
	}
	return tag.Name, nil
}

func checkLatestVersionFromAPI() (string, error) {
	// TODO
	return "", errors.New("not support now")
}

func ShowVersion(f factory.Factory) {
	log := f.GetLog()
	if common.Version == "" {
		common.Version = defaultVersion
	}
	if common.BuildDate == "" {
		common.BuildDate = defaultBuildDate
	}
	if common.GitCommitHash == "" {
		common.GitCommitHash = defaultGitCommitHash
	}
	tmpl, err := newVersionTemplate()
	if err != nil {
		log.Fatalf("gen version failed, reason: %v", err)
		return
	}
	vd := versionInfo{
		Client: clientVersion{
			Version:      common.Version,
			GoVersion:    runtime.Version(),
			GitCommit:    common.GitCommitHash,
			BuildTime:    common.BuildDate,
			Os:           runtime.GOOS,
			Arch:         runtime.GOARCH,
			Experimental: true,
		},
	}
	log.Debug("check update...")
	lastVersion, lastType, err := PreCheckLatestVersion(log)
	if err != nil {
		log.Debugf("get update message err: %v", err)
	}
	if lastVersion != "" && !strings.Contains(common.Version, lastVersion) {
		nowVersion := gv.MustParse(strings.TrimPrefix(common.Version, "v"))
		needUpgrade := nowVersion.LessThan(gv.MustParse(lastVersion))
		if needUpgrade {
			vd.Client.CanUpgrade = true
			vd.Client.LastVersion = lastVersion
			vd.Client.Version = color.SGreen(vd.Client.Version)
			vd.Client.UpgradeMessage = fmt.Sprintf("Now you can use %s to upgrade cli to the latest version %s by %s mode", color.SGreen("%s upgrade", os.Args[0]), color.SGreen(lastVersion), color.SGreen(lastType))
		}
	}
	if err := prettyPrintVersion(vd, tmpl); err != nil {
		panic(err)
	}
}

func prettyPrintVersion(vd versionInfo, tmpl *template.Template) error {
	t := tabwriter.NewWriter(os.Stdout, 20, 1, 1, ' ', 0)
	err := tmpl.Execute(t, vd)
	t.Write([]byte("\n"))
	t.Flush()
	return err
}

func newVersionTemplate() (*template.Template, error) {
	tmpl, err := template.New("version").Parse(versionTpl)
	return tmpl, errors.Wrap(err, "template parsing error")
}
