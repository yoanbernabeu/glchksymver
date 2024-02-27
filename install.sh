#!/usr/bin/env bash

echo "Installing glchksymver..."
echo "------------------------"

# Determining the Linux distribution and architecture
distro=$(lsb_release -i -s)
arch=$(uname -m)

echo "Distribution: $distro"
echo "Architecture: $arch"

# glchksymver version
version="v0.1.0"

echo "Version: $version"
echo "------------------------"

# URL for downloading the archive based on the distribution and architecture
url=""

case "$distro" in
  "Darwin")
    case "$arch" in
      "x86_64")
        url="https://github.com/yoanbernabeu/glchksymver/releases/download/${version}/glchksymver-${version}-darwin-amd64.tar.gz"
        ;;
      "arm64")
        url="https://github.com/yoanbernabeu/glchksymver/releases/download/${version}/glchksymver-${version}-darwin-arm64.tar.gz"
        ;;
      *)
        echo "Unsupported architecture"
        exit 1
        ;;
    esac
    ;;
  "Ubuntu"|"Debian"|"Raspbian")
  echo "Downloading glchksymver..."
    case "$arch" in
      "i686")
        url="https://github.com/yoanbernabeu/glchksymver/releases/download/${version}/glchksymver-${version}-linux-386.tar.gz"
        ;;
      "x86_64")
        url="https://github.com/yoanbernabeu/glchksymver/releases/download/${version}/glchksymver-${version}-linux-amd64.tar.gz"
        echo $url
        ;;
      "arm64")
        url="https://github.com/yoanbernabeu/glchksymver/releases/download/${version}/glchksymver-${version}-linux-arm64.tar.gz"
        ;;
      *)
        echo "Unsupported architecture"
        exit 1
        ;;
    esac
    ;;
  *)
    echo "Unsupported distribution"
    exit 1
    ;;
esac

# Downloading the archive to home directory (and check if url is not 404)
echo "Downloading glchksymver..."
wget -q --spider $url
if [ $? -eq 0 ]; then
  wget -O ~/glchksymver.tar.gz $url -q --show-progress
else
  echo "------------------------"
  echo "glchksymver archive not found"
  echo "------------------------"
  exit 1
fi

# Extracting the archive (if it exists)
echo "Extracting glchksymver..."
if [ -f ~/glchksymver.tar.gz ]; then
  tar -xzf ~/glchksymver.tar.gz -C ~/
else
  echo "glchksymver archive not found"
  exit 1
fi

# Removing the archive
echo "Removing archive..."
rm ~/glchksymver.tar.gz

# Moving the binary to /usr/local/bin
echo "Moving glchksymver to /usr/local/bin..."
sudo mv ~/glchksymver /usr/local/bin/

# Making the binary executable
echo "Making glchksymver executable..."
sudo chmod +x /usr/local/bin/glchksymver

# Sending a message to the user
echo "-----------------------------------------"
echo "glchksymver successfully installed"
echo "-----------------------------------------"