class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.1.3"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "46b95a93121ddc6def120ca66395cd9bb3651462182d3e956c6c8b9c5279f030"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "b179c6223cb2d7ed4b3f9769ea5dbfccf98dee96d792f36c57e63b4b1a0ec104"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "9c6da0ef54719c9736e83489b5cd9b3437bda7ec27b550dbf3b8f9b0e08da0c7"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "6944709c1849b83b66d98a3cd32839813d71282a642d99c94914fb8378d38cc8"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
