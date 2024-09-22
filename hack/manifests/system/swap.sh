#!/bin/bash

set -e

swapsize=${1:-"1G"}
swappath=${2:-"/swapfile"}

# apt install util-linux
fallocate -l "$swapsize" "$swappath"

chmod 600 "$swappath"

mkswap "$swappath"

swapon "$swappath"

cp -a /etc/fstab /etc/fstab.bak

echo "$swappath swap swap defaults 0 0" >> /etc/fstab

free -m
