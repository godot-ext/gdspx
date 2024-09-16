#!/bin/bash

# Default values
EXPORT_WEB=false
PROJECT_PATH="tutorial/01_aircraft"

# Parse command line arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        --export) EXPORT_WEB=true ;;
        --path) PROJECT_PATH="$2"; shift ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

# Get current directory name
CURRENT_PATH=$("$PWD")
mkdir -p ./builds/web_go

# Build wasm
cd "$PROJECT_PATH/pkg" || { echo "Project path not found"; exit 1; }
echo "Building Go wasm"
GOOS=js GOARCH=wasm go build -o "$CURRENT_PATH/builds/web_go/gdspx.wasm"
cd $CURRENT_PATH

# Export to web if enabled
if [ "$EXPORT_WEB" = true ]; then
    echo "================ EXPORT_WEB ====================="
    ./godot/bin/godot.linuxbsd.editor.dev.x86_64 --headless --quit --path "./$PROJECT_PATH/project" --export-debug "Web" "$CURRENT_PATH/builds/web_go/index.html"
fi
