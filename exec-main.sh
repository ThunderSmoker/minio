#!/bin/bash

# Check if command-line arguments are provided
if [ $# -eq 0 ]; then
  echo "No command provided. Exiting..."
  exit 1
fi


go run main.go
