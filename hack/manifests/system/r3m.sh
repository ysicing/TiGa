#!/bin/bash

DOWNLOADER=

info()
{
    echo '[INFO] ' "$@"
}

fatal()
{
    echo '[ERROR] ' "$@" >&2
    exit 1
}

verify_downloader() {
    # Return failure if it doesn't exist or is no executable
    [ -x "$(command -v $1)" ] || return 1

    # Set verified executable as our downloader program and return success
    DOWNLOADER=$1
    return 0
}

verify_downloader wget || verify_downloader curl ||  fatal 'Can not find curl or wget for downloading files'
set +e
case $DOWNLOADER in
  curl)
    curl -o /usr/bin/tiga-r3m -sfL https://c.ysicing.net/oss/tiga/linux/amd64/r3m
    ;;
  wget)
    wget -qO /usr/bin/tiga-r3m https://c.ysicing.net/oss/tiga/linux/amd64/r3m
    ;;
  *)
      fatal "Incorrect downloader executable '$DOWNLOADER'"
    ;;
esac
[ $? -eq 0 ] || fatal 'Download failed'
set -e

chmod +x /usr/bin/tiga-r3m

info "r3m install success"

mkdir -p /etc/tiga

cat > /etc/tiga/r3m.toml <<EOF
[log]
level = "warn"
output = "stdout"

[network]
no_tcp = false
use_udp = true
tcp_timeout = 5
udp_timeout = 30

[[endpoints]]
listen = "0.0.0.0:5000"
remote = "1.1.1.1:443"

[[endpoints]]
listen = "0.0.0.0:5001"
remote = "1.1.1.1:443"
extra_remotes = ["1.1.1.2:443", "1.1.1.3:443"]
balance = "roundrobin: 4, 2, 1"
EOF

cat > /etc/systemd/system/r3m.service <<EOF
[Unit]
Description=tiga-r3m
After=network-online.target
Wants=network-online.target systemd-networkd-wait-online.service

[Service]
Type=simple
User=root
Restart=on-failure
RestartSec=5s
DynamicUser=true
WorkingDirectory=/etc/tiga
ExecStart=/usr/bin/tiga-r3m -c /etc/tiga/r3m.toml

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable r3m --now

info "r3m install done"
