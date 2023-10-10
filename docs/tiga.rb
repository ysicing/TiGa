class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.9"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "68177163de72cdbca63afd4afd77c14551dad01c9927fb85e513657847554fd5"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "1d03bba8efea6aec0ea07cd9167d42cb6c01b90d16aa1353f12a35f5b201ab40"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "19f631553e6aaf975ea0fed37414d81e627a273deb3d5876a5acfa487ab5c757"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "29fcaa40a4ed100403274d16ac22d8d28b2b7ca0433053a3f64f570bcba3e25a"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
