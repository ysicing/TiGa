class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.13"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "0107be96d4520dd10813c97fe7ef2ef5cd1f7428ae0a48cc0ae68ba5f079b287"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "b66aed25cfc331faf9c3f2fdb6791618f614c42b00b007aa3b9739302102d241"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "d1439adf9fce026bcf541144a545c269499fc2d21857d7b8555cc2334f0921e0"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "5e30ebcf612c466293a4d223463486c54fd76883e87b8cc38795a84531163a44"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
