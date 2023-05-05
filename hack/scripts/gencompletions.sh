#!/usr/bin/env bash

set -ex

rm -rf completions
mkdir completions
go build -o /tmp/tiga
for sh in bash zsh fish; do
	/tmp/tiga completion "$sh" >"completions/tiga.$sh"
done
rm -rf /tmp/tiga
