# TiGA

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ysicing/tiga?filename=go.mod&style=flat-square)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/ysicing/tiga?style=flat-square)
![GitHub](https://img.shields.io/badge/license-YPL%20%2B%20AGPL-blue)
![GitHub all releases](https://img.shields.io/github/downloads/ysicing/tiga/total?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/ysicing/tiga)](https://goreportcard.com/report/github.com/ysicing/tiga)
[![Releases](https://img.shields.io/github/release-pre/ysicing/tiga.svg)](https://github.com/ysicing/ergo/releases)

## Introduction

TiGA is a devops toolset, including a set of tools for daily development and operation and maintenance.

## Install

### macOS

```bash
# 暂不支持 brew install tiga
brew install ysicing/tap/tiga
```

### Debian/Ubuntu

```bash
echo "deb [trusted=yes] https://mirrors.ysicing.cloud/ysicing/apt/ /" | tee /etc/apt/sources.list.d/ysicing.list
apt update
apt install tiga
```

### CentOS

```bash
cat > /etc/yum.repos.d/ysicing.repo << EOF
[ysicing]
name=Quickon Repo
baseurl=https://mirrors.ysicing.cloud/ysicing/yum/
enabled=1
gpgcheck=0
EOF

yum makecache
yum install tiga
```

## Build

```bash
git clone https://github.com/ysicing/tiga.git
cd tiga
# install task https://taskfile.dev/#/installation
go install github.com/go-task/task/v3/cmd/task@latest
# build
task -v
```

## Contributors

<!-- readme: collaborators,contributors -start -->
<!-- readme: collaborators,contributors -end -->
<a href="https://github.com/easysoft/quickon_cli/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=ysicing/tiga" />
</a>
