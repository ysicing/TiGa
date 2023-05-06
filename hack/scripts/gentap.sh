#!/usr/bin/env bash

version=$(cat VERSION)
macosamd64sha=$(cat dist/checksums.txt | grep tiga_darwin_amd64 | awk '{print $1}')
macosarm64sha=$(cat dist/checksums.txt | grep tiga_darwin_arm64| awk '{print $1}')
linuxamd64sha=$(cat dist/checksums.txt | grep tiga_linux_amd64 | awk '{print $1}')
linuxarm64sha=$(cat dist/checksums.txt | grep tiga_linux_arm64 | awk '{print $1}')

cat > docs/tiga.rb <<EOF
class Tiga < Formula
    desc "Simple and powerful tool for sernior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "${version}"

    if OS.mac?
      if Hardware::CPU.arm?
        url "https://ghproxy.com/https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "${macosarm64sha}"
      else
        url "https://ghproxy.com/https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "${macosamd64sha}"
      end
    elsif OS.linux?
      if Hardware::CPU.intel?
        url "https://ghproxy.com/https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "${linuxamd64sha}"
      end
      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://ghproxy.com/https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "${linuxarm64sha}"
      end
    end

    def install
      if OS.mac?
        if Hardware::CPU.intel?
          bin.install "tiga_darwin_amd64" => "tiga"
        else
          bin.install "tiga_darwin_arm64" => "tiga"
        end
      elsif OS.linux?
        if Hardware::CPU.intel?
          bin.install "tiga_linux_amd64" => "tiga"
        else
          bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end

    test do
      assert_match "tiga vervion v#{version}", shell_output("#{bin}/tiga version")
    end
end
EOF
