#!/bin/bash

set -e

swapsize=${1:-"1G"}

# apt install util-linux
fallocate -l $swapsize /swapfile

chmod 600 /swapfile

mkswap /swapfile

swapon /swapfile

cp -a /etc/fstab /etc/fstab.bak

echo "/swapfile swap swap defaults 0 0" >> /etc/fstab

free -m
