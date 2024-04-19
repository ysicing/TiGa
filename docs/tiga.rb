class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.3.12"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "d017261e04a9f9a6e775ac33df507624c813e6d48af827fed2f411cbe159d342"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "0433e49f6f6f996e759dc4c28f4d570d370c91566a4daec6d1e356a9b00e3cb1"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "0d501a2cb833976f579f9b25168acf7f0cec85689f1eb5b25be34869272e9811"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "c4d37b56424f3b03f654e9346877773c0de8170894fa3f2a4ef52c3a194fbdcf"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
