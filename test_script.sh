#!/bin/bash

GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m'

folder_path="examples"  # Specify the path to the folder containing the files

# Use a for loop to iterate over files in the specified folder
for filename in "$folder_path"/*; do
    if [ -f "$filename" ]; then
        echo -e "$GREEN------ Running $filename ------$NC"
        echo
        go run . "$filename"
        echo -en "\n${YELLOW}Press enter to continue...${NC}"
        read
        echo
    fi
done
