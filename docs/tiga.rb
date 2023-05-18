class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.0.14"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "cca4bea55f82304faaa738a9b80bf825335f5243cf5b3c89d04c08193ca0bab2"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "87ccd0f0ad9d287e2e6b0fd0543b52af8c835757ece522fa4e2a632d9ea32b4d"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "0677cbaf66a0d8f720ae5d99d76ba6628d2d061c93afa7050c0a8882395a113b"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "bca9cfc55fb31dff4435ba41671b96f891263616a7bdcb94a4c0fb71228e383e"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
