class Tiga < Formula
    desc "Simple and powerful tool for sernior restart engineer"
    homepage "https://github.com/ysicing/tiga"
    version "0.0.7"

    if OS.mac?
      if Hardware::CPU.arm?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_arm64"
        sha256 "745cbc099a52279ff65db7f3a9a9056eab19d4ccf5c51852bf14c1c3a48daec0"
      else
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_darwin_amd64"
        sha256 "7ca995f7bfb440661057b08c8aead6f543e385dcab2f7090af6763837eea8854"
      end
    elsif OS.linux?
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_amd64"
        sha256 "935198d9ec266940f9333db5e78b2754f8386ca33f3e057965f18606e28bf4cc"
      end
      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/ysicing/tiga/releases/download/v#{version}/tiga_linux_arm64"
        sha256 "f604dab942e77d2bff98bc46fca37fd9e6334ddd23a720485f7f14ad4368ae29"
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
