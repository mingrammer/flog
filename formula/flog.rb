class Flog < Formula
  desc "A fake log generator for common log formats"
  homepage "https://github.com/mingrammer/flog"
  url "https://github.com/mingrammer/flog/releases/download/v0.2.1/flog_0.2.1_darwin_amd64.tar.gz"
  version "0.2.1"
  sha256 "9d2efabe6a6a51b1ad1acfbd132a9668620a6dd28a44fc93983a5f92e48e5efd"

  def install
    bin.install "flog"
  end
end
