#!/bin/bash

echo "Building git_comment..."
go build -o git_comment ./cmd/git_comment_main
if [ $? -eq 0 ]; then
    echo "✓ Build successful: ./git_comment"
else
    echo "✗ Build failed"
    exit 1
fi
