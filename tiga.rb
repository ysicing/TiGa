class Tiga < Formula
    desc "Simple and powerful tool for sernior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.0.6"

    if OS.mac?
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "8b51bb02fe01bd12622bb5c7b23b4b426077d9f2b5d7e2849ffdb69241767394"
      else
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "c7d4365f198839c0397efffb23d6272db21730552064a35fe6c468d0f0d8cf00"
      end
    elsif OS.linux?
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "ff00a52aa611bb8aebac83eb69401b721adacdf6b11a7929ac311276ac22892e"
      end
      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "5d0c738ad11e60e102a4a15ebd701555e892db9aa33fa77c4f95bffa7c74cd69"
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
