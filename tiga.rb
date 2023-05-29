class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.0.15"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "e7376a1a15d6fb5d515ae5a162ff2919228c6e52050972607e3bdfedf7beddf1"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "0201a085a197e9b0722e99434ceda475b6cfa9fd79649a05ccb29b4c793407da"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "8ecbc70aef7b5f4c3336376a2414ac192c353aef5e7762226f4ef72f9af70670"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "cc9644d9387cb4550d65a4b759cb26bfd94a16a8726b639ea4a221cd632cb3ce"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
