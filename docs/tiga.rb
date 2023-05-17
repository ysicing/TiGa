class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.0.13"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "b7d51ccc58bd1b8069c9a8adb0bb91626ffab6ea9c67df5585655ffc26964a1b"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "c996e5e556da037c4553798ff9d8b2bad938f7074332532bfb282886a8dace11"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "2225c4d461c713e6759f8a6ed10dc1f0e56494aca33383e8f46066b818ec4562"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "04980c9b9fbb87633d9fe07c55f510cdfd24f7fa64a0d0d3f7dba58816d9c4a7"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
