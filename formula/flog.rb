class Flog < Formula
  desc "A fake log generator for common log formats"
  homepage "https://github.com/mingrammer/flog"
  url "https://github.com/mingrammer/flog/releases/download/v0.1.4/flog_0.1.4_darwin_amd64.tar.gz"
  version "0.1.4"
  sha256 "fc857e591e8d9b18fe6a3c55d374a4b4f84407d42a3dac88eadec943c90acbfb"

  def install
    bin.install "flog"
  end
end
