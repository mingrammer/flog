class Flog < Formula
  desc "A fake log generator for common log formats"
  homepage "https://github.com/mingrammer/flog"
  url "https://github.com/mingrammer/flog/releases/download/v0.1.0/flog_0.1.0_darwin_amd64.tar.gz"
  version "0.1.0"
  sha256 "3d98d2de4a9a5a5cf458c0dbfb570338b53d24cd7445b75b1fc5aebe59ca748a"

  def install
    bin.install "flog"
  end
end
