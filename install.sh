#!/bin/bash

CDN_URL="https://cdn.gofred.io"
TEMP_DIR="/tmp"
OUTPUT_DIR="$HOME/.local/bin"

# Check env if gofred is already installed
if [ -f "gofred" ]; then
    echo "gofred is already installed. If you want to update it, run 'gofred update'"
    exit 0
fi

# Get OS and ARCH
OS=$(uname -s)
ARCH=$(uname -m)

if [ $OS == "Darwin" ]; then
    OUTPUT_DIR="$HOME/.local/bin"
    if [ $ARCH == "arm64" ]; then
        OS="darwin-arm64"
    else
        OS="darwin-amd64"
    fi
fi

if [ $OS == "Linux" ]; then
    OUTPUT_DIR="$HOME/.local/bin"
    if [ $ARCH == "arm64" ]; then
        OS="linux-arm64"
    else
        OS="linux-amd64"
    fi
fi

if [ $OS == "Windows" ]; then
    OUTPUT_DIR="$HOME/AppData/Local/bin"
    if [ $ARCH == "arm64" ]; then
        OS="windows-arm64"
    else
        OS="windows-amd64"
    fi
fi

# Create the output directory if it doesn't exist
if [ ! -d $OUTPUT_DIR ]; then
    mkdir -p $OUTPUT_DIR
fi

# Download the index.json file
curl -L -o $TEMP_DIR/gofred.index.json $CDN_URL/cli/index.json

# Parse the index.json file
VERSION=$(cat $TEMP_DIR/gofred.index.json | jq -r '.version')

# Download the gofred binary
curl -L -o $OUTPUT_DIR/gofred $CDN_URL/cli/$VERSION/gofred-$OS

# Make the file executable if OS is not Windows
if [ $OS != "windows-arm64" ] && [ $OS != "windows-amd64" ]; then
    chmod +x $OUTPUT_DIR/gofred
fi

# Make sure OUTPUT_DIR is in the PATH
if ! echo $PATH | grep -q $OUTPUT_DIR; then
    echo "$OUTPUT_DIR is not in the PATH. Do you want to add it? (y/n)"
    read addToPath
    if [ $addToPath == "y" ]; then
        # Detect the shell
        if [ $SHELL == "/bin/bash" ]; then
            echo "export PATH=$PATH:$OUTPUT_DIR" >> ~/.bashrc
            source ~/.bashrc
        elif [ $SHELL == "/bin/zsh" ]; then
            echo "export PATH=$PATH:$OUTPUT_DIR" >> ~/.zshrc
            source ~/.zshrc
        fi
    fi
fi

# Print the path to the gofred binary
echo "gofred installed to $OUTPUT_DIR/gofred"

# Print the version of gofred
echo "gofred version: $(gofred version)"

# Print success message
echo "gofred installed successfully"
echo "Run 'gofred --help' to get started"