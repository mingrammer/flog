class Flog < Formula
  desc "A fake log generator for common log formats"
  homepage "https://github.com/mingrammer/flog"
  url "https://github.com/mingrammer/flog/releases/download/v0.0.5/flog_0.0.5_darwin_amd64.tar.gz"
  version "0.0.5"
  sha256 "d4e55b1202da41fb1d57d32959b3695c2cd14931d3bf8c2af7e9bf02cac4be9c"

  def install
    bin.install "flog"
  end
end
