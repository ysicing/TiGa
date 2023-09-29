class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.5"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "95837aeca474fe9c8bbd336fb7b36c9be0012d66a33ac1cc186e8e47067e28fb"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "e9c6ff1cf8f0e8176ce23012a9a8da688f0eed6143e1dc041418fcf7a24311ce"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "411bd069b7055f0f4423530f7bd405d73ceff6906f7d02676b4cf9b081b2bb93"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "5ba6263b68311d935ca1593472b865233eddfaa446d5588f92529f4bf5b6cb96"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
