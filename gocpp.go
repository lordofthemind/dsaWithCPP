package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Function to extract the GCC version
func getGCCVersion() (string, error) {
	cmd := exec.Command("g++", "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(output), "\n")
	if len(lines) > 0 {
		fields := strings.Fields(lines[0])
		if len(fields) >= 4 {
			return fields[3], nil
		}
	}
	return "", fmt.Errorf("could not determine GCC version")
}

// Function to compile and run the C++ program
func compileAndRun(cppFile, cppStandard string, debugMode, noClean bool) error {
	// Extract the filename without extension
	filename := strings.TrimSuffix(cppFile, ".cpp")

	// Compile the program
	compileArgs := []string{"-std=" + cppStandard, "-Wall", "-Wextra", "-o", filename, cppFile}
	if debugMode {
		compileArgs = append(compileArgs, "-g")
	}

	compileCmd := exec.Command("g++", compileArgs...)
	compileCmd.Stdout = os.Stdout
	compileCmd.Stderr = os.Stderr

	fmt.Println("🔨 Compiling with standard optimizations...")
	if err := compileCmd.Run(); err != nil {
		return fmt.Errorf("❌ Compilation failed: %v", err)
	}
	fmt.Println("✅ Compilation successful!")

	// Run the compiled program
	fmt.Println("🚀 Running", filename, "...")
	runCmd := exec.Command("./" + filename)
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr

	if err := runCmd.Run(); err != nil {
		return fmt.Errorf("❌ Program execution failed: %v", err)
	}

	// Clean up the compiled file unless noClean is true
	if !noClean {
		fmt.Println("🧹 Cleaning up...")
		if err := os.Remove(filename); err != nil {
			return fmt.Errorf("❌ Failed to clean up: %v", err)
		}
	}

	fmt.Println("✨ All done!")
	return nil
}

func main() {
	// Default settings
	cppStandard := "c++20"
	debugMode := false
	noClean := false

	// Parse command-line arguments
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("❌ Usage: gcpp <cpp_file> [-s c++17|c++20] [-d] [-n]")
		os.Exit(1)
	}

	cppFile := args[0]
	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "-s":
			if i+1 < len(args) && (args[i+1] == "c++17" || args[i+1] == "c++20") {
				cppStandard = args[i+1]
				i++
			} else {
				fmt.Println("❌ Error: Invalid C++ standard. Use 'c++17' or 'c++20'.")
				os.Exit(1)
			}
		case "-d":
			debugMode = true
		case "-n":
			noClean = true
		default:
			fmt.Println("❌ Error: Unrecognized option", args[i])
			os.Exit(1)
		}
	}

	// Check if the source file exists
	if _, err := os.Stat(cppFile); os.IsNotExist(err) {
		fmt.Println("❌ Error: File", cppFile, "not found!")
		os.Exit(1)
	}

	// Display compiler information
	gccVersion, err := getGCCVersion()
	if err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}

	fmt.Println("------------------------------------------------------------")
	fmt.Println("🔧    Compiler GCC:")
	fmt.Println("🛠️     GCC Version:", gccVersion)
	fmt.Println("📋    C++ Standard Version:", cppStandard)
	fmt.Println("🚩    Compiler Flags: -std=" + cppStandard + " -Wall -Wextra")
	fmt.Println("------------------------------------------------------------")

	// Compile and run the program
	startTime := time.Now()
	if err := compileAndRun(cppFile, cppStandard, debugMode, noClean); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("⏱️  Total execution time: %v\n", time.Since(startTime))
}