#!/bin/bash

echo "Running comprehensive golangci-lint check..."

# Use the -c flag to specify the 'all linters' config file
# Run on the entire project recursively ('./...')
golangci-lint run -c .golangci-all.yml ./...

# Check the exit code
if [ $? -ne 0 ]; then
  echo "golangci-lint found issues."
  exit 1
fi

echo "golangci-lint check passed."
exit 0
