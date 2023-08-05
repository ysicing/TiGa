class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.2.3"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "13ac030367e58e1162a8c6dec35a725782d836e91763e1d0ff96894fe48766c8"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "94524f61cca43b4419e79f54b903dd35fbe9d91f7fe34ed6ab874fcfb1c3c2f0"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "3969f73863bdbcff2471cdc619c1e44b9009d3128de3dac9dd7d9c0188131e93"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "82a51764952bf30f53a618689fe84cb82d96ab9c3819cf0e22b41e86b065d39a"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
