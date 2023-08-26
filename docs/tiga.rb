class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.0"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "c5063b381b03b85706bb84413d1fdd244daf35e36669bbf56390efa921c8ebb5"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "2c255ba8f3dd8c5268872d660dcfa91f28d76cc2fd8bd93c54877758f11c2a22"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "bb04a3a7de84d2b5f78ca45a77975d545f7b1e2ed7d9c6fae9f853212d0eab87"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "0dad088eef45476953dd63eaba667f5e5f79ef3822121ea21c65b43a88f09020"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
