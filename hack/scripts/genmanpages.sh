#!/usr/bin/env bash

set -e

rm -rf manpages
mkdir manpages
go run ./main.go man | gzip -c -9 >manpages/tiga.1.gz
