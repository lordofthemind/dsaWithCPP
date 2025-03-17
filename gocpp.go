package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Color codes for terminal output
const (
	colorReset   = "\033[0m"
	colorRed     = "\033[38;5;196m"
	colorOrange  = "\033[38;5;208m"
	colorYellow  = "\033[38;5;214m"
	colorGreen   = "\033[38;5;46m"
	colorBlue    = "\033[38;5;27m"
	colorCyan    = "\033[38;5;51m"
	colorPurple  = "\033[38;5;105m"
	colorPink    = "\033[38;5;45m"
	colorWhite   = "\033[38;5;15m"
	colorGrey    = "\033[38;5;245m"
	colorTeal    = "\033[38;5;87m"
	colorMagenta = "\033[38;5;201m"
)

// Configuration options
type Config struct {
	CPPStandard string
	DebugMode   bool
	VerboseMode bool
	NoClean     bool
	CPPFile     string
}

// Print a horizontal line with specified color
func printLine(color string) {
	line := strings.Repeat("▁", 80)
	fmt.Printf("%s%s%s\n", color, line, colorReset)
	line = strings.Repeat("▔", 80)
	fmt.Printf("%s%s%s\n", color, line, colorReset)
}

// Print a styled message
func printMessage(color, message string) {
	printLine(colorCyan)
	fmt.Printf("%s%s%s\n", color, message, colorReset)
	printLine(colorGreen)
}

// Show error and exit
func showError(message string) {
	printLine(colorCyan)
	fmt.Printf("%s❌ Error: %s%s\n", colorRed, message, colorReset)
	printLine(colorGreen)
	os.Exit(1)
}

// Get GCC version
func getGCCVersion() string {
	cmd := exec.Command("g++", "--version")
	output, err := cmd.Output()
	if err != nil {
		return "Unknown"
	}
	
	// Parse the output to extract version
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 4 {
			return parts[3]
		}
	}
	
	return "Unknown"
}

// Show compiler information
func showCompilerInfo(config *Config) {
	printLine(colorCyan)
	fmt.Printf("%s🔧    Compiler GCC:%s\n", colorBlue, colorReset)
	fmt.Printf("%s🛠️     GCC Version: %s%s\n", colorPink, getGCCVersion(), colorReset)
	fmt.Printf("%s📋    C++ Standard Version: %s%s\n", colorCyan, config.CPPStandard, colorReset)
	fmt.Printf("%s🚩    Compiler Flags: -std=%s -Wall -Wextra%s\n", colorTeal, config.CPPStandard, colorReset)
	printLine(colorCyan)
}

// Execute command and print output
func executeCommand(command string, args ...string) (bool, error) {
	cmd := exec.Command(command, args...)
	
	// Set up pipes for stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return false, err
	}
	
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return false, err
	}
	
	// Start the command
	if err := cmd.Start(); err != nil {
		return false, err
	}
	
	// Copy stdout and stderr to console
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)
	
	// Wait for the command to finish
	err = cmd.Wait()
	return err == nil, err
}

// Main function
func main() {
	// Check if G++ is installed
	_, err := exec.LookPath("g++")
	if err != nil {
		fmt.Println("Error: G++ compiler is not installed. Please install it first.")
		os.Exit(1)
	}
	
	// Set up configuration
	config := Config{
		CPPStandard: "c++20", // Default C++ standard
	}
	
	// Set up command line flags
	flag.BoolVar(&config.DebugMode, "d", false, "Enable debug mode")
	flag.BoolVar(&config.VerboseMode, "v", false, "Enable verbose mode")
	flag.BoolVar(&config.NoClean, "n", false, "Disable cleanup after execution")
	standardPtr := flag.String("s", "c++20", "C++ standard (c++17 or c++20)")
	
	// Parse flags
	flag.Parse()
	
	// Validate C++ standard
	if *standardPtr != "c++17" && *standardPtr != "c++20" {
		showError("Invalid C++ standard. Use 'c++17' or 'c++20'.")
	}
	config.CPPStandard = *standardPtr
	
	// Get the source file
	if flag.NArg() < 1 {
		printMessage(colorOrange, "❌ Usage: "+os.Args[0]+" <cpp_file> [-s c++17|c++20] [-d] [-v] [-n]")
		os.Exit(1)
	}
	
	config.CPPFile = flag.Arg(0)
	
	// Check if the source file exists
	if _, err := os.Stat(config.CPPFile); os.IsNotExist(err) {
		showError(fmt.Sprintf("File '%s' not found!", config.CPPFile))
	}
	
	// Extract filename without extension
	filename := strings.TrimSuffix(filepath.Base(config.CPPFile), filepath.Ext(config.CPPFile))
	
	// Enable verbose mode if requested
	if config.VerboseMode {
		printMessage(colorOrange, "🔍 Verbose mode enabled - tracing commands...")
	}
	
	// Display compiler information
	showCompilerInfo(&config)
	
	// Compile the program with appropriate flags
	compileArgs := []string{"-std=" + config.CPPStandard, "-Wall", "-Wextra"}
	
	if config.DebugMode {
		printMessage(colorOrange, "🐞 Debug mode enabled!")
		compileArgs = append(compileArgs, "-g")
	} else {
		printMessage(colorCyan, "🔨 Compiling with standard optimizations...")
	}
	
	compileArgs = append(compileArgs, "-o", filename, config.CPPFile)
	
	// Print compile command if verbose
	if config.VerboseMode {
		fmt.Println("Executing: g++", strings.Join(compileArgs, " "))
	}
	
	// Compile
	success, _ := executeCommand("g++", compileArgs...)
	
	// Check if compilation was successful
	if success {
		printMessage(colorGreen, "✅ Compilation successful!")
		fmt.Printf("%s🚀 Running %s...%s\n", colorBlue, filename, colorReset)
		printLine(colorCyan)
		
		// Program output section
		printLine(colorCyan)
		fmt.Printf("%s📤 Program Output:%s\n", colorBlue, colorReset)
		printLine(colorCyan)
		
		// Run the program
		programPath := "./" + filename
		if runtime.GOOS == "windows" {
			programPath = filename + ".exe"
		}
		
		success, _ := executeCommand(programPath)
		printLine(colorCyan)
		
		if !success {
			printMessage(colorRed, "❌ Program execution failed!")
		}
	} else {
		showError("Compilation failed!")
	}
	
	// Clean up the compiled file unless -n is passed
	if !config.NoClean {
		printMessage(colorPurple, "🧹 Cleaning up...")
		err := os.Remove(filename)
		if err != nil && config.VerboseMode {
			fmt.Printf("Warning: Could not remove file: %s\n", err)
		}
		printMessage(colorGreen, "✨ All done!")
	} else {
		printMessage(colorYellow, fmt.Sprintf("📝 Keeping binary file: %s", filename))
		printMessage(colorGreen, "✨ All done!")
	}
}