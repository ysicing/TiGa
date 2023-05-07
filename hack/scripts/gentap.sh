#!/usr/bin/env bash

version=$(cat VERSION)
# shellcheck disable=SC2002
macosAMD64sha=$(cat dist/checksums.txt | grep tiga_darwin_amd64 | awk '{print $1}')
# shellcheck disable=SC2002
macosARM64sha=$(cat dist/checksums.txt | grep tiga_darwin_arm64| awk '{print $1}')
# shellcheck disable=SC2002
linuxAMD64sha=$(cat dist/checksums.txt | grep tiga_linux_amd64 | awk '{print $1}')
# shellcheck disable=SC2002
linuxARM64sha=$(cat dist/checksums.txt | grep tiga_linux_arm64 | awk '{print $1}')

cat > docs/tiga.rb <<EOF
class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "${version}"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "${macosARM64sha}"

        def install
            bin.install "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "${macosAMD64sha}"

        def install
            bin.install "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "${linuxAMD64sha}"

        def install
            bin.install "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "${linuxARM64sha}"

        def install
            bin.install "tiga"
        end
      end
    end
end
EOF
