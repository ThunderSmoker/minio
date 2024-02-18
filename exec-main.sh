#!/bin/bash

# Check if command-line arguments are provided
if [ $# -eq 0 ]; then
  echo "No command provided. Exiting..."
  exit 1
fi

# Run main.go with the provided command-line arguments
cd /minio-project
go run main.go
