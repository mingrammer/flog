class Flog < Formula
  desc "A fake log generator for common log formats"
  homepage "https://github.com/mingrammer/flog"
  url "https://github.com/mingrammer/flog/releases/download/v0.1.3/flog_0.1.3_darwin_amd64.tar.gz"
  version "0.1.3"
  sha256 "c65d94cefed801f00d51d91bea867107ed65bdc7f09d6c1456911e12aecf99bd"

  def install
    bin.install "flog"
  end
end
