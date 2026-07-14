#!/bin/bash

echo "Building CommitSmith..."
go build -o cmtr ./cmd/cmtr
if [ $? -eq 0 ]; then
    echo "✓ Build successful: ./cmtr"
else
    echo "✗ Build failed"
    exit 1
fi
