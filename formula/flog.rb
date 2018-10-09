class Flog < Formula
  desc "A fake log generator for common log formats"
  homepage "https://github.com/mingrammer/flog"
  url "https://github.com/mingrammer/flog/releases/download/v0.2.0/flog_0.2.0_darwin_amd64.tar.gz"
  version "0.2.0"
  sha256 "78938f3462d5ce79a7ad603ac8ae15be4fc705d1232ed0ec532b6f939348b19d"

  def install
    bin.install "flog"
  end
end
