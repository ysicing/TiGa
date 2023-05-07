#!/usr/bin/env bash

set -eu

bin=$1

echo "upx compress tiga ${bin}"

command_exists() {
  command -v "$@" > /dev/null 2>&1
}

if command_exists upx; then
  if [ -f "${bin}" ]; then
      upx --ultra-brute "${bin}"
  else
      echo "not found ${bin}"
  fi
fi
