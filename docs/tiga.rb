class Tiga < Formula
    desc "Simple and powerful tool for sernior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.0.5"

    if OS.mac?
      if Hardware::CPU.arm?
        url "https://ghproxy.com/https://github.com/ysicing/tiga/releases/download/#{version}/tiga_darwin_arm64"
        sha256 "6cb5eccb075b23b72b74143f301725ea81b6a983fa158ce26fb4a47008d8ffa7"
      else
        url "https://ghproxy.com/https://github.com/ysicing/tiga/releases/download/#{version}/tiga_darwin_amd64"
        sha256 "a72c72c1147b427a65134ddaa41334eb900044c49f03585398c3bfb30fd48266"
      end
    elsif OS.linux?
      if Hardware::CPU.intel?
        url "https://ghproxy.com/https://github.com/ysicing/tiga/releases/download/#{version}/tiga_linux_amd64"
        sha256 "27fbebb8d058b98a9dec7dbec2e55f538370fa4cb563850561ce759319a0c01e"
      end
      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://ghproxy.com/https://github.com/ysicing/tiga/releases/download/#{version}/tiga_linux_arm64"
        sha256 "f99ef5dd15dfa8b3923c216e4353885ed67452a74088daaa8ee2da706bc8b440"
      end
    end

    def install
      if OS.mac?
        if Hardware::CPU.intel?
          bin.install "tiga_darwin_amd64" => "tiga"
        else
          bin.install "tiga_darwin_arm64" => "tiga"
        end
      elsif OS.linux?
        if Hardware::CPU.intel?
          bin.install "tiga_linux_amd64" => "tiga"
        else
          bin.install "tiga_linux_arm64" => "tiga"
        end
      end
    end

    test do
      assert_match "tiga vervion v#{version}", shell_output("#{bin}/tiga version")
    end
end
