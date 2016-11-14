require 'formula'

class Ddget < Formula
  VERSION = '0.1.1'

  homepage 'https://github.com/winebarrel/ddget'
  url "https://github.com/winebarrel/ddget/releases/download/v#{VERSION}/ddget-v#{VERSION}-darwin-amd64.gz"
  sha256 'd5dc1c8f58d2bb6d5590c3c0cf877544f32851a58614c27c93b2be929c9fe126'
  version VERSION
  head 'https://github.com/winebarrel/ddget.git', :branch => 'master'

  def install
    system "mv ddget-v#{VERSION}-darwin-amd64 ddget"
    bin.install 'ddget'
  end
end
