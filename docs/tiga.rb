class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.0.8"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "ebfcece17e5209c8025a791f30c0dd4ec629bf20bcae89a03444dc482f54a381"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "aa4d58c5982fbb6887f5a3b8f4a7f4595128b205829b2b6c1d06ddbaf6cfceb6"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "3e893a71aa4f1b161b2aba800dcb4ff24da2341d38f04984f7660b8854728903"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "2c285fb534202b7659045cfb22f78a92e46d4d605e4e35728c9b228a77ff4a5d"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
