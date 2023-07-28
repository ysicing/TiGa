class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.2.0"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "d4146b7e564cf349ee9cef3a6484febd5ca5ba630544c2c8f8c5f407220e02cc"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "c0844371d16cd731f5889e8090b7698d85fed93678849276c7aad2a88d274358"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "f87451b9ff6ab5dffe3a75eb6ef57d772361a5f418df480877e3fd05b9bdd03c"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "2de2d606c88b7f7373a83b68240f8e520a624f29e3f4c4a56776cbd4a0d1480e"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
