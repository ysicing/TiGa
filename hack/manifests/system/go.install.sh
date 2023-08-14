#!/bin/bash

set -e

gofile=$1

[ -f "$gofile" ] || exit 1

rm -rf /usr/local/go && tar -C /usr/local -xzf "$gofile"

rm -rf "$gofile"
