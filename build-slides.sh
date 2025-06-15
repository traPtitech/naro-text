#!/bin/bash

set -eux

# Create dist directory if it doesn't exist
mkdir -p slides/dist

# Find all slide.md files in slides subdirectories and build them
for slide_file in $(find slides -name "slide.md" -type f); do
    # Get the directory name containing the slide.md
    dir_name=$(basename "$(dirname "$slide_file")")
    
    # Build with marp and output to dist folder
    echo "Building $slide_file -> slides/dist/${dir_name}.pdf"
    npx marp --theme slides/style.css --allow-local-files "$slide_file" -o "slides/dist/${dir_name}.pdf"
done

echo "All slides built successfully!"