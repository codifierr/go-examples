#!/bin/bash

# Get the parent folder path
parent_folder="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Function to recursively run 'go get -u ./...' and 'go mod tidy' in module folders
function process_module {
    local module_dir="$1"

    # Check if the directory is a Go module
    if [[ -f "$module_dir/go.mod" ]]; then
        echo "Processing module in '$module_dir'"

        # Run 'go get -u ./...' to update dependencies
        (cd "$module_dir" && go get -u ./...)

        # Run 'go mod tidy' to clean up the module
        (cd "$module_dir" && go mod tidy)
    fi

    # Recursively process subdirectories
    for dir in "$module_dir"/*; do
        if [[ -d "$dir" ]]; then
            process_module "$dir"
        fi
    done
}

# Main script
process_module "$parent_folder"