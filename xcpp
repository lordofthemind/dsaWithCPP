#!/bin/bash

# Function to extract the GCC version
get_gcc_version() {
    g++ --version | head -n 1 | awk '{print $4}'
}

# Default settings
cpp_standard="c++20"  # Default C++ standard version
debug_mode=false       # Debug mode disabled by default
verbose_mode=false     # Verbose mode disabled by default
no_clean=false         # Cleanup enabled by default

# Function to display compiler information
show_compiler_info() {
    echo "------------------------------------------------------------"
    echo "🔧    Compiler GCC:"
    echo "🛠️     GCC Version: $(get_gcc_version)"
    echo "📋    C++ Standard Version: $cpp_standard"
    echo "🚩    Compiler Flags: -std=$cpp_standard -Wall -Wextra"
    echo "------------------------------------------------------------"
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
                echo "❌ Error: Invalid C++ standard. Use 'c++17' or 'c++20'."
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
                echo "❌ Error: Unrecognized option '$1'."
                exit 1
            fi
            ;;
    esac
done

# Ensure a C++ source file is provided
if [ -z "$cpp_file" ]; then
    echo "❌ Usage: $0 <cpp_file> [-s c++17|c++20] [-d] [-v] [-n]"
    exit 1
fi

# Extract the filename without extension
filename=$(basename "$cpp_file" .cpp)

# Check if the source file exists
if [ ! -f "$cpp_file" ]; then
    echo "❌ Error: File '$cpp_file' not found!"
    exit 1
fi

# Enable verbose mode if requested
if [ "$verbose_mode" = true ]; then
    set -x  # Enable command tracing for debugging
fi

# Display compiler information
show_compiler_info

# Compile the program with appropriate flags
if [ "$debug_mode" = true ]; then
    echo "🐞 Debug mode enabled!"
    g++ -std=$cpp_standard -Wall -Wextra -g -o "$filename" "$cpp_file"
else
    g++ -std=$cpp_standard -Wall -Wextra -o "$filename" "$cpp_file"
fi

# Check if compilation was successful
if [ $? -eq 0 ]; then
    echo "✅ Compilation successful!"
    echo "🚀 Running $filename..."
    echo "------------------------------------------------------------"
    ./"$filename"
    echo "------------------------------------------------------------"
else
    echo "❌ Compilation failed!"
    exit 1
fi

# Clean up the compiled file unless -n is passed
if [ "$no_clean" = false ]; then
    echo "🧹 Cleaning up..."
    rm "$filename"
fi
