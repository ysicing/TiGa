class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.2"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "21cc4c6458127ff2a33469ebeafb8bacdfca2a85c5be85164be7d73212743758"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "03411c204803d329572ec0cb1ba759b09716d6381adfadfd992d54f1a47ae824"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "1e75e6cc5df86ae0a2f76e27fd1d2136e26a746534693a4e53e9be08a0676f6e"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "ca02d27b3c26f5bc23ed70d068396d01de55a412f8a2b38a5dcd520af2a0e98a"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
