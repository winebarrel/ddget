require 'formula'

class Ddget < Formula
  VERSION = '0.1.0'

  homepage 'https://github.com/winebarrel/ddget'
  url "https://github.com/winebarrel/ddget/releases/download/v#{VERSION}/ddget-v#{VERSION}-darwin-amd64.gz"
  sha256 '05f8c34b40d67c3cb25befff6acb762120c34dc60448657280dfa55fa0e469cb'
  version VERSION
  head 'https://github.com/winebarrel/ddget.git', :branch => 'master'

  def install
    system "mv ddget-v#{VERSION}-darwin-amd64 ddget"
    bin.install 'ddget'
  end
end
