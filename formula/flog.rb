class Flog < Formula
  desc "A fake log generator for common log formats"
  homepage "https://github.com/mingrammer/flog"
  url "https://github.com/mingrammer/flog/releases/download/v0.0.4/flog_0.0.4_darwin_amd64.tar.gz"
  version "0.0.4"
  sha256 "7dd94d901412445baf1e88c208f382a91575c952dcfb257dec8f8e6e9ce18150"

  def install
    bin.install "flog"
  end
end
