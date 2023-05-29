class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.1.1"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "621aa48a88fc2bd3e0443ca4460159d2aa14f1359d6e8c75d900fc636d9b9780"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "8d3ea1cdfb0e5e7b7432b006d00a9d6916f4de43bae502a436a6a93d263db19d"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "335ebb2d4e67958449f24f6c5a69e9611aeb6556060ddbb6d4ceb9875a2e7e89"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "8b31e9ec70ce73a556dd4d9426579935d1f79f5eb643667d5690ed524be1b82b"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
