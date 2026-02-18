#
# Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
# Date: 2026-02-17
#

# Homebrew formula for the Tradier CLI. Downloads pre-built binaries
# from GitHub releases for the current platform and architecture.
class Tradier < Formula
  desc "CLI tool for the Tradier brokerage API"
  homepage "https://github.com/cloudmanic/tradier"
  license "MIT"
  version "latest"

  if OS.mac? && Hardware::CPU.arm?
    url "https://github.com/cloudmanic/tradier/releases/latest/download/tradier-darwin-arm64"
  elsif OS.mac? && Hardware::CPU.intel?
    url "https://github.com/cloudmanic/tradier/releases/latest/download/tradier-darwin-amd64"
  elsif OS.linux? && Hardware::CPU.arm?
    url "https://github.com/cloudmanic/tradier/releases/latest/download/tradier-linux-arm64"
  elsif OS.linux? && Hardware::CPU.intel?
    url "https://github.com/cloudmanic/tradier/releases/latest/download/tradier-linux-amd64"
  end

  def install
    binary_name = stable.url.split("/").last
    bin.install binary_name => "tradier"
  end

  test do
    assert_match "tradier version", shell_output("#{bin}/tradier --version")
  end
end
