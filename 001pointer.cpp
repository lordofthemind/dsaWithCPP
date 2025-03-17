#include <iostream>
using namespace std;

int main()
{
    int a = 5, *ptr;
    ptr = &a;
    cout << "a=" << a << endl;
    cout << "a=" << *ptr << endl;
    return 0;
}
