class Flog < Formula
  desc "A fake log generator for common log formats"
  homepage "https://github.com/mingrammer/flog"
  url "https://github.com/mingrammer/flog/releases/download/v0.0.1/flog_0.0.1_darwin_amd64.tar.gz"
  version "0.0.1"
  sha256 "e2a6fb756494147bc25afef5421bcb8dcd737002ff4eb332de0b92d871514a59"

  def install
    bin.install "flog"
  end
end
