class TiGA < Formula
    desc "Simple and powerful tool for sernior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.0.5"

    if OS.mac?
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/#{version}/tiga_darwin_arm64"
        sha256 "f1795d17ee7a273c84a12a9ca1be16d010f3fe6c4474f6dc483ab3fe45d54ec5"
      else
        url "https://github.com/ysicing/tiga/releases/download/#{version}/tiga_darwin_amd64"
        sha256 "d046e21f8932727337481efa55cd726a8d7c8c8120959e422580d667667aede2"
      end
    elsif OS.linux?
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/#{version}/tiga_linux_amd64"
        sha256 "b1ec94c42123b5fd34a9413f000dbe88e017f613e90e94eb83fc2ac12af1bb22"
      end
      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/#{version}/tiga_linux_arm64"
        sha256 "98a955517ef298727fee6120e976c2c0cbd4d3ea70b3e6ad8d19834e8d2fa260"
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
