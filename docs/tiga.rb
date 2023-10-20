class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.10"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "5aac88ffad3a7d362682e885784ac6e4b61cacb4b8ae3b77d4a871cc0f58a567"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "0c3874f6576512ffeed587900e04e8ca1dbbf23ef8aed81c6d7b56ef973ef658"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "95db0fba4d00e2102d1247c6d0460206a1c98fe4d8e0ac5125ed18854815fa9f"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "d158ec4e2045d3410197e7844ab5548571dab3970f5de4fd27be502c5b180d36"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
