#include <iostream>
#include <string>
#include <array>
#include <memory>
#include <stdexcept>
#include <cstdio>
#include <algorithm>

// Function to execute a shell command and return the output
std::string exec(const char *cmd)
{
    std::array<char, 128> buffer;
    std::string result;
    std::unique_ptr<FILE, decltype(&pclose)> pipe(popen(cmd, "r"), pclose);
    if (!pipe)
    {
        throw std::runtime_error("popen() failed!");
    }
    while (fgets(buffer.data(), buffer.size(), pipe.get()) != nullptr)
    {
        result += buffer.data();
    }
    return result;
}

// Helper function to trim whitespace from the end of a string
std::string trim(const std::string &str)
{
    auto end = std::find_if_not(str.rbegin(), str.rend(), [](unsigned char ch)
                                { return std::isspace(ch); });
    return std::string(str.begin(), end.base());
}

int main()
{
    // Check if Gum is installed
    try
    {
        std::string gum_version = exec("gum --version");
        std::cout << "Gum is installed. Version: " << gum_version;
    }
    catch (const std::runtime_error &e)
    {
        std::cerr << "❌ Gum CLI is not installed. Please install it first.\n";
        std::cerr << "Visit: https://github.com/charmbracelet/gum\n";
        return 1;
    }

    // Use Gum for a styled heading
    std::string header = exec("gum style --foreground 212 --bold --border double 'C++ with Gum Demo'");
    std::cout << header;

    // Use Gum to get user input
    std::cout << "\nLet's get some user input using Gum:\n";
    std::string name = exec("gum input --placeholder 'Enter your name'");
    name = trim(name); // Trim any trailing newline
    std::cout << "You entered: " << name << "\n";

    // Use Gum to create a styled message
    std::string message = "echo 'Hello, " + name + "!' | gum style --foreground 99 --italic";
    std::cout << exec(message.c_str());

    // Use Gum for a confirmation
    std::cout << "\nConfirm exit:\n";
    std::string confirm = exec("gum confirm 'Do you want to exit?' && echo 'Yes' || echo 'No'");
    std::cout << "You chose: " << confirm << "\n";

    return 0;
}