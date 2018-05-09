#include <stdio.h>
#include "test.h"


void hello() {
 printf("Hello, World!\n");
}
#ifdef __TEST__ // 避免和 Go bootstrap main 冲突。 这里也可以单独编译链接后运行此c程序

int main(int argc, char *argv[]) {
 hello();
 return 0;
}

#endif