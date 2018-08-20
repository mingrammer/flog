class Flog < Formula
  desc "A fake log generator for common log formats"
  homepage "https://github.com/mingrammer/flog"
  url "https://github.com/mingrammer/flog/releases/download/v0.1.1/flog_0.1.1_darwin_amd64.tar.gz"
  version "0.1.1"
  sha256 "715ca38b6f9e5156b5e18385c00bf3aecbfe3d1130bb87f96c17cd1812f32abc"

  def install
    bin.install "flog"
  end
end
