class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.3"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "f66b6ec151f4236ffa25bd5176bc72edc8d15bc47a918b8e8aada638d887d8b0"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "553a1df6be881b924fbc7c79e3df6aa5996c9ca671371b0e610bc96b42ac9ed1"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "73c1cdc2e64790b8bd017d48f87002d7d9c95f5d55ccb278b2b3a1327528d047"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "d127e979362de375eaa2200321d0f84f646ec9679db99760914360ba271caf31"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
