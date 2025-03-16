#include <iostream>

// This will allow us to check the C++ standard version
#if __cplusplus >= 202002L
#define CPP_VERSION "C++20 or newer"
#elif __cplusplus >= 201703L
#define CPP_VERSION "C++17"
#elif __cplusplus >= 201402L
#define CPP_VERSION "C++14"
#elif __cplusplus >= 201103L
#define CPP_VERSION "C++11"
#elif __cplusplus >= 199711L
#define CPP_VERSION "C++98/C++03"
#else
#define CPP_VERSION "Pre-standard C++"
#endif

int main()
{
    std::cout << "✅ C++ is successfully installed!" << std::endl;

    // Display the C++ version
    std::cout << "📋 C++ Standard Version: " << CPP_VERSION << std::endl;
    std::cout << "📊 Raw __cplusplus value: " << __cplusplus << std::endl;

// Display compiler information
#if defined(__GNUC__) && !defined(__clang__)
    std::cout << "🔧 Compiler: GCC" << std::endl;
    std::cout << "🔄 Version: " << __GNUC__ << "." << __GNUC_MINOR__ << "." << __GNUC_PATCHLEVEL__ << std::endl;
#elif defined(__clang__)
    std::cout << "🔧 Compiler: Clang" << std::endl;
    std::cout << "🔄 Version: " << __clang_major__ << "." << __clang_minor__ << "." << __clang_patchlevel__ << std::endl;
#elif defined(_MSC_VER)
    std::cout << "🔧 Compiler: Microsoft Visual C++" << std::endl;
    std::cout << "🔄 Version: " << _MSC_VER << std::endl;
#else
    std::cout << "🔧 Compiler: Unknown" << std::endl;
#endif

    // Check if the program successfully runs
    std::cout << "\n✨ If you can see this message, your C++ compiler is working correctly!" << std::endl;

    return 0;
}