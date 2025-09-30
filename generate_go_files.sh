#!/bin/bash

# Check if directory argument is provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <directory_path>"
    echo "Example: $0 ./icons"
    exit 1
fi

DIRECTORY="$1"
EMBEDS_FILE="embeds.go"
RESOURCES_FILE="resources.go"

# Check if directory exists
if [ ! -d "$DIRECTORY" ]; then
    echo "Error: Directory '$DIRECTORY' does not exist"
    exit 1
fi

# Initialize output files if they don't exist
if [ ! -f "$EMBEDS_FILE" ]; then
    echo "package main" > "$EMBEDS_FILE"
    echo "" >> "$EMBEDS_FILE"
fi

if [ ! -f "$RESOURCES_FILE" ]; then
    echo "package main" > "$RESOURCES_FILE"
    echo "" >> "$RESOURCES_FILE"
    echo "import (" >> "$RESOURCES_FILE"
    echo "    \"fyne.io/fyne/v2/storage/repository\"" >> "$RESOURCES_FILE"
    echo "    \"your-module/assets\"" >> "$RESOURCES_FILE"
    echo ")" >> "$RESOURCES_FILE"
    echo "" >> "$RESOURCES_FILE"
fi

# Function to check if a variable already exists in a file
variable_exists() {
    local var_name="$1"
    local file="$2"
    grep -q "var ${var_name}SVG \[\]byte" "$file" 2>/dev/null
}

resource_exists() {
    local var_name="$1"
    local file="$2"
    grep -q "var ${var_name}icon = fyne.NewStaticResource" "$file" 2>/dev/null
}

# Counter for new files processed
new_count=0
skipped_count=0

# Process each SVG file in the directory
for filepath in "$DIRECTORY"/*.svg; do
    # Check if any SVG files exist
    if [ ! -e "$filepath" ]; then
        echo "No SVG files found in '$DIRECTORY'"
        exit 1
    fi
    
    # Extract just the filename
    filename=$(basename "$filepath")
    
    # Remove the .svg extension for variable name
    var_name="${filename%.svg}"
    
    # Check if this variable already exists in both files
    if variable_exists "$var_name" "$EMBEDS_FILE" && resource_exists "$var_name" "$RESOURCES_FILE"; then
        echo "Skipping $filename (already processed)"
        skipped_count=$((skipped_count + 1))
        continue
    fi
    
    # Add to embeds file if not already there
    if ! variable_exists "$var_name" "$EMBEDS_FILE"; then
        echo "//go:embed icons/$filename" >> "$EMBEDS_FILE"
        echo "var ${var_name}SVG []byte" >> "$EMBEDS_FILE"
        echo "" >> "$EMBEDS_FILE"
    fi
    
    # Add to resources file if not already there
    if ! resource_exists "$var_name" "$RESOURCES_FILE"; then
        echo "var ${var_name}icon = fyne.NewStaticResource(\"$filename\", ${var_name}SVG)" >> "$RESOURCES_FILE"
        echo "" >> "$RESOURCES_FILE"
    fi
    
    echo "Added $filename -> ${var_name}SVG & ${var_name}icon"
    new_count=$((new_count + 1))
done

echo ""
echo "Summary:"
echo "  - New files processed: $new_count"
echo "  - Files skipped (already exist): $skipped_count"
echo "  - Output files: $EMBEDS_FILE, $RESOURCES_FILE"


#./generate_go_files.sh /path/to/your/svg/directory