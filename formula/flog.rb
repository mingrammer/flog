class Flog < Formula
  desc "A fake log generator for common log formats"
  homepage "https://github.com/mingrammer/flog"
  url "https://github.com/mingrammer/flog/releases/download/v0.1.2/flog_0.1.2_darwin_amd64.tar.gz"
  version "0.1.2"
  sha256 "bebbf043a0831edfe03712269fa235621653c759f89d8696dfc28035259ba27a"

  def install
    bin.install "flog"
  end
end
