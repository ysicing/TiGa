class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.4"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "c41e2c6a485281b1615399aba5fafb9956693dab637aa89eed3fe6f74ed3a974"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "6f966b846c301e51c79b41c9b09f7181cf7876244822884f0a4960706106f6f6"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "c98612626780b932c3dd164d24a08830e6d8131e0ccb35c651d733fc4a61b4df"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "a4ae5578cdea759d7778ad1e789793966de5d34f3a07f3df140c9fb05504ebef"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
