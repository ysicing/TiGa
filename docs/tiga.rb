class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.0.10"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "ffb7401642e2c1a070d4cb014c9b864ae964b837a22ebfa4e53ba215df563d05"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "b8482883b0328cc2aa714aee628553630d37e200346ad6a1ed3af957be37c523"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "1af891d790ef4695607f57527e17e0292195f24b6b56a61b738f697d5ba4c45a"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "a3878c06eabf5bba384afa2d7aabc0e58b8b22a416bb19f0069e4d2a37003c23"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
