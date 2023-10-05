class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.7"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "ed9a3d10c102430cc865e78fabb1785d268e5a7191b7aec821c61511dfc84d47"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "b445170886797c2a0385075494931fc960c4bab4ae850f24f31770ef33c7b13f"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "515e73927cb0443394529fab5e5a8e3d17fd1c69ec87223129a95a8e571310f0"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "5115d1dbeb42a7b5cc5eb85c90739f6afc98dbd8315b61b7bd28a50ccb3b6747"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
