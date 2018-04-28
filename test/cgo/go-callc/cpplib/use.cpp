// use_mylib.cpp
#include <iostream>
#include <string>
using namespace std;
#ifdef __cplusplus
extern "C" {
#endif
#include "mylib.h" // for foobar()
#ifdef __cplusplus
}
#endif
int main()
{
  foobar(100);
  return 0;
}