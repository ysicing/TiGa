class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.0.9"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "25a1d108ed88985ad9a6f6b5dcdfb1cb0fd9fabf6fefea0a3ee5881cbce50e05"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "1b6eec5709f9a0937816c10c03ab1f81b2f686deb7079bb39b8e093c66f20c8a"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "db341e1c7150fc5df7c1d58c080451bae9bd72f2f5bddf8402a9a610fd01f6fc"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "797da963f404fc4e464089c41e32c71d4a53f3826a990bc6bb5d2198dcb81291"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
