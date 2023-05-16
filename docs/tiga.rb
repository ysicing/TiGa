class Tiga < Formula
    desc "Simple and powerful tool for senior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.0.12"

    on_macos do
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "e8ecd4840b57e72d77406eb8db1ef8b9cc40ff337b3c2d51885014887b595e20"

        def install
            bin.install "tiga_darwin_arm64" => "tiga"
        end
      end

      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "fc232a32c9ac421ccba84711b99af3fd405e923690d65383b8d8e493483d5e4d"

        def install
            bin.install "tiga_darwin_amd64" => "tiga"
        end
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "4a47e4bebbf2763006edefaea3da52eea207f5f203a863d10975a8fe13fe0b02"

        def install
            bin.install "tiga_linux_amd64" => "tiga"
        end
      end

      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "3a87e536735cb8e074d85bc4a86b2c460e994e1d05e1e0856ddd808dd04da0d9"

        def install
            bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end
end
