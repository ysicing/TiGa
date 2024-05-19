class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.15"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "3ef7089e2c4ea1e77c2b41d8286096ba26b8cdba2c98fb910ca5bd527eebc918"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "ea930687bffd198a25b7e29b83792302ee3a2dd9533bcb5453434303761a510a"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "993493c579a468a169770a736fa69d23a958dc6dc1a9fb9c5231e8b20699eccc"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "0ac75f90079e372bb9550af4103f26b92f30570824151a13a3d58215936c6c69"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
