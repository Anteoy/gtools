#cpplib
gcc -fPIC -shared -o libmy.so mylib.c
g++  use.cpp -L. -lmy
export LD_LIBRARY_PATH=./
./a.out




