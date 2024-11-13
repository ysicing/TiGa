class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.16"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "0b979cbe574bcc35cc982cffa6eb78793cbc6037af97158a731205d4065410b3"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "9336f029de40cd2108d6c53c5d559c1b8df7eb217b3a6a71fb4075e2a5fc514d"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "f266e7e0e8e2afa8d30bc5fda79b60311008c4fc04ba3d51f2adb859fc6f6d1d"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "5058a72c256792cd7095ee521bbb54d1438a1f540bea4f31536b4421688f4f6c"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
