class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.14"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "f589eacc2e8b44ed9ba76243882d1ef5f627c0c2d840b3465bffb9612e9ca8f3"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "efc393417c1da5258ba48286b5c45eeb25cd8db03d3c6e475116ee535fec9c63"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "2dc605136b0abb01a7e18d95458b302c4b99fdf265c431d0a7f94d5c9003a4b9"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "a666c6042d9a0b9f221af8051cb820fbdd380b34097153dec22f1668a2af0337"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
