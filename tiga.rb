class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.1.2"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "1a9f2b1f30a01fc7421512d4a6eeccc767abff9b3fe420d65045832b42323b8d"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "e911981ca7f23533b6339e620be7c85c0a68ffd48791315c14739db7d2181168"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "cb0cecdf2d7fffff57656e6bf0c2c87f115cc522412ed722206e9ad63590b712"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "adfb31f6cca5fd8d4fd8979eb6c0cd13af231c148efd3490ac251c601a200ea7"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
