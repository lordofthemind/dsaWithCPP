#!/bin/bash

# Function to extract the GCC version
get_gcc_version() {
    g++ --version | head -n 1 | awk '{print $4}'
}

# Default settings
cpp_standard="c++20"  # Default C++ standard version
debug_mode=false      # Debug mode disabled by default
verbose_mode=false    # Verbose mode disabled by default
no_clean=false        # Cleanup enabled by default

# Function to display compiler information
show_compiler_info() {
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
    gum style --foreground 39 "🔧    Compiler GCC:"
    gum style --foreground 45 "🛠️     GCC Version: $(get_gcc_version)"
    gum style --foreground 51 "📋    C++ Standard Version: $cpp_standard"
    gum style --foreground 87 "🚩    Compiler Flags: -std=$cpp_standard -Wall -Wextra"
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
}

# Parse command-line options
while [[ $# -gt 0 ]]; do
    case "$1" in
        # Set C++ standard (only supports c++17 or c++20)
        -s)
            if [[ "$2" == "c++17" || "$2" == "c++20" ]]; then
                cpp_standard="$2"
                shift 2  # Move past the argument
            else
                gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
                gum style --foreground 196 "❌ Error: Invalid C++ standard. Use 'c++17' or 'c++20'."
                gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
                exit 1
            fi
            ;;
        # Enable debug mode
        -d) debug_mode=true; shift ;;
        # Enable verbose mode (for debugging script execution)
        -v) verbose_mode=true; shift ;;
        # Disable cleaning after execution
        -n) no_clean=true; shift ;;
        # Handle the input file argument
        *)
            if [[ -z "$cpp_file" ]]; then
                cpp_file="$1"
                shift
            else
                gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
                gum style --foreground 196 "❌ Error: Unrecognized option '$1'."
                gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
                exit 1
            fi
            ;;
    esac
done

# Ensure a C++ source file is provided
if [ -z "$cpp_file" ]; then
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 202 "❌ Usage: $0 <cpp_file> [-s c++17|c++20] [-d] [-v] [-n]"
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
    exit 1
fi

# Extract the filename without extension
filename=$(basename "$cpp_file" .cpp)

# Check if the source file exists
if [ ! -f "$cpp_file" ]; then
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 196 "❌ Error: File '$cpp_file' not found!"
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
    exit 1
fi

# Enable verbose mode if requested
if [ "$verbose_mode" = true ]; then
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 208 "🔍 Verbose mode enabled - tracing commands..."
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
    set -x  # Enable command tracing for debugging
fi

# Display compiler information
show_compiler_info

# Compile the program with appropriate flags
if [ "$debug_mode" = true ]; then
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 208 "🐞 Debug mode enabled!"
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
    g++ -std=$cpp_standard -Wall -Wextra -g -o "$filename" "$cpp_file"
else
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 51 "🔨 Compiling with standard optimizations..."
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
    g++ -std=$cpp_standard -Wall -Wextra -o "$filename" "$cpp_file"
fi

# Check if compilation was successful
if [ $? -eq 0 ]; then
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 46 "✅ Compilation successful!"
    gum style --foreground 27 "🚀 Running $filename..."
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
    
    # Program output section
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 39 "📤 Program Output:"
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
    ./"$filename"
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
else
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 196 "❌ Compilation failed!"
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
    exit 1
fi

# Clean up the compiled file unless -n is passed
if [ "$no_clean" = false ]; then
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 105 "🧹 Cleaning up..."
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
    rm "$filename"
    
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 46 "✨ All done!"
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
else
    gum style --foreground 33 "$(gum join --horizontal $(printf "▁%.0s" {1..80}))"
    gum style --foreground 214 "📝 Keeping binary file: $filename"
    gum style --foreground 46 "$(gum join --horizontal $(printf "▔%.0s" {1..80}))"
fi