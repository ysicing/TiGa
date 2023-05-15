class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.0.11"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "2f14368e3e1ba53d25f65acb426b8af5d1caf3b2c4de7dba75dcc55fbe317d69"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "adabe7cb33c2109739cd3979714f24a7953e19c3e556a74d40cf8b9d205f1103"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "8adba0212ed619bd5926c6f5dcb906944fd85d5f9cf5c4941cb2b82149865175"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "304bc7a9b222b37df6044ceb9a14b690e541e058f64872aacc4d19352bc3c64d"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
