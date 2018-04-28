#ifndef __MY_LIB_H__
#define __MY_LIB_H__
void foobar(int i);
#endif /* __MY_LIB_H__ */

// gcc -fPIC -shared -o libmy.so mylib.c

// 2这里的共享对象其实全称是动态共享对象文件（Dynamic Shared Objects，简写为DSO）；

// 2-fPIC：地址无关代码（Position-Independent Code），该技术主要用于解决SO中对绝对地址的重定位问题；