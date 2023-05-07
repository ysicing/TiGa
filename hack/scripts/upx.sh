#!/usr/bin/env bash

set -eu

os=$1
arch=$2

echo "upx compress tiga ${os} ${arch}"

command_exists() {
	command -v "$@" > /dev/null 2>&1
}

if command_exists upx; then
  upx --ultra-brute dist/tiga_"${os}"_"${arch}"/tiga
fi
