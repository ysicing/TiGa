class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.8"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "9bdfa89f98e6f40a95c4db7db610fa779b32637c1bb970ee21a17ee98d6ab493"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "69d5294395db664218035a4d6b3796b3222f02d327e9755f8bfcaba90338eb1f"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "a07d32eab51feb1e83e5334e6d54681a78b9453d1c986c2349ed8fb4cad5afdc"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "03bcfd407d705d21b473d8978673ab8ec873954453d1fb29d84e820767dc7fc3"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
