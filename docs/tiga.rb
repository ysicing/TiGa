class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.17"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "f13c0962220ab693d301b6f0bbfeb58d0ef14702bd625382307904926def0f1f"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "4f147f26e66577cb51ad9c1e4106ed02d7718d833ff9fc4f6a6304f267fba312"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "cfefebcb31890c97df3deacddc391599b2980b8d065783139a49a98d2c46b0df"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "0ea970dc19d324c393707e9d33feba20c70d3273fea5b4b47b90e5acda5e214e"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
