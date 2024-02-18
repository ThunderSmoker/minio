#!/bin/bash

# Check if command-line arguments are provided
if [ $# -eq 0 ]; then
  echo "No command provided. Exiting..."
  exit 1
fi

# Check the first argument provided
if [ "$1" == "ls" ]; then
  go run main.go
elif [ "$1" == "rm" ]; then
  # Check if filename is provided as the second argument
  if [ -z "$2" ]; then
    echo "Filename not provided with 'rm' command. Exiting..."
    exit 1
  fi
  go run main.go "$2"
else
  echo "Unknown command: $1"
  exit 1
fi
