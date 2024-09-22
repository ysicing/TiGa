#!/bin/bash

name=$1
home=$2
binurl=$3

[ -z "$binurl" ] && (
    echo "binary url is empty"
    exit 1
)

# download binary
wget -O $home/.tiga/bin/tiga-$name $binurl

chmod +x $home/.tiga/bin/tiga-$name
