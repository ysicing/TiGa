class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.2.1"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "3dc76e5cc6e251dd8e7784b871c091de965a80ee29e43813ab6847d998e73915"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "069c5940097ba0610c37ba3ab9aa827e5ada010b8a2224a054092c1250f3a25b"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "9e018cb41cabd286e1ea458961808b4ceb51c5f9cf042a42da0181a53a40c7e1"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "3c3258fd1ea62fd68a9fdf4648d7429270bb06ac493a55353705294d8a77402f"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
